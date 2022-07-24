package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"hijrastep/graph/generated"
	"hijrastep/graph/model"
	"hijrastep/graph/util"
	"hijrastep/models"
)

func (r *contentResolver) IsCompleted(ctx context.Context, obj *models.Content) (bool, error) {
	if len(obj.Logs) > 0 {
		return true, nil
	}

	return false, nil
}

func (r *mutationResolver) CreateBab(ctx context.Context, input model.NewBab) (*models.Bab, error) {
	bab := models.Bab{
		Title:   input.Title,
		Order:   input.Order,
		SubBabs: []models.SubBab{},
	}

	for _, subab := range input.SubBabs {
		bab.SubBabs = append(bab.SubBabs, models.SubBab{
			Title:          subab.Title,
			Order:          subab.Order,
			ImageURL:       subab.ImageURL,
			InstructorName: subab.InstructorName,
		})
	}

	result := r.DB.Save(&bab)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &bab, nil
}

func (r *mutationResolver) CreateSubBab(ctx context.Context, babID int, input model.NewSubBab) (*models.SubBab, error) {
	subBab := models.SubBab{
		BabID:          babID,
		Title:          input.Title,
		Order:          input.Order,
		ImageURL:       input.ImageURL,
		InstructorName: input.InstructorName,
	}

	result := r.DB.Create(&subBab)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &subBab, nil
}

func (r *mutationResolver) CreateMaterial(ctx context.Context, subBabID int, input model.NewMaterial) (*models.Material, error) {
	material := models.Material{
		VideoURL: input.VideoURL,
		AudioURL: input.AudioURL,
		Text:     input.Text,
		Content: models.Content{
			SubBabID:         subBabID,
			Title:            input.Title,
			Order:            input.Order,
			Type:             "material",
			DurationInMinute: input.DurationInMinute,
		},
	}

	result := r.DB.Create(&material)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &material, nil
}

func (r *mutationResolver) CreateQuiz(ctx context.Context, subBabID int, input model.NewTest) (*models.Test, error) {
	quiz := models.Test{
		MinimumScore: input.MinimumScore,
		Content: models.Content{
			SubBabID: subBabID,
			Title:    input.Title,
			Order:    input.Order,
			Type:     "quiz",
		},
	}

	result := r.DB.Create(&quiz)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &quiz, nil
}

func (r *mutationResolver) CreateExam(ctx context.Context, subBabID int, input model.NewTest) (*models.Test, error) {
	exam := models.Test{
		MinimumScore: input.MinimumScore,
		Content: models.Content{
			SubBabID: subBabID,
			Title:    input.Title,
			Order:    input.Order,
			Type:     "exam",
		},
	}

	result := r.DB.Create(&exam)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &exam, nil
}

func (r *mutationResolver) CreateQuestion(ctx context.Context, testID int, input model.NewQuestion) (*models.Question, error) {
	question := models.Question{
		TestID:   testID,
		Order:    input.Order,
		Question: input.Question,
		Options:  input.Options,
		Hint:     input.Hint,
		Answer:   input.Answer,
	}

	result := r.DB.Create(&question)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &question, nil
}

func (r *mutationResolver) MarkComplete(ctx context.Context, contentID int, email string) (bool, error) {
	log := models.ContentLog{
		ContentID: contentID,
		Email:     email,
	}

	result := r.DB.Create(&log)

	if result.Error != nil {
		return false, fmt.Errorf("%v", result.Error)
	}

	return true, nil
}

func (r *mutationResolver) SubmitExam(ctx context.Context, testID int, email string, answers []int) (*models.ExamLog, error) {
	var test models.Test
	result := r.DB.Preload("Content.Logs", "Email = ?", email).Preload("Questions", util.FuncOrder).Preload("Logs", "Email = ?", email).Find(&test, testID)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	int32slice := []int32{}
	correctAnswerCounter := 0

	// Hitung jumlah benar sekalian convert []int jadi []int32
	for i := 0; i < len(answers); i++ {
		if answers[i] == test.Questions[i].Answer {
			correctAnswerCounter++
		}

		int32slice = append(int32slice, int32(answers[i]))
	}

	score := correctAnswerCounter * 100 / len(test.Questions)

	examLog := models.ExamLog{
		TestID:           testID,
		Email:            email,
		Score:            score,
		CorrectAnswerNum: correctAnswerCounter,
		Answers:          int32slice,
	}

	// Check if log already exist (pernah ngerjain sebelumnya)
	if len(test.Logs) > 0 {

		// Update kalau nilainya lebih besar
		if score > test.Logs[0].Score {
			test.Logs[0].Score = score
			test.Logs[0].CorrectAnswerNum = correctAnswerCounter
			test.Logs[0].Answers = int32slice

			r.DB.Save(&test.Logs[0])
		}
	} else {
		result := r.DB.Create(&examLog)

		if result.Error != nil {
			return nil, fmt.Errorf("%v", result.Error)
		}
	}

	// Mark sebagai done kalau nilainya melebihi kkm
	if len(test.Content.Logs) == 0 && score >= test.MinimumScore {
		log := models.ContentLog{
			ContentID: testID,
			Email:     email,
		}

		result := r.DB.Create(&log)

		if result.Error != nil {
			return nil, fmt.Errorf("%v", result.Error)
		}
	}

	return &examLog, nil
}

func (r *mutationResolver) DeleteBab(ctx context.Context, babID int) (bool, error) {
	result := r.DB.Delete(&models.Bab{}, babID)

	if result.Error != nil {
		return false, fmt.Errorf("%v", result.Error)
	}

	return true, nil
}

func (r *mutationResolver) DeleteSubBab(ctx context.Context, subBabID int) (bool, error) {
	result := r.DB.Delete(&models.SubBab{}, subBabID)

	if result.Error != nil {
		return false, fmt.Errorf("%v", result.Error)
	}

	return true, nil
}

func (r *queryResolver) Babs(ctx context.Context) ([]*models.Bab, error) {
	var babs []*models.Bab
	result := r.DB.Preload("SubBabs", util.FuncOrder).Order("\"order\"").Find(&babs)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return babs, nil
}

func (r *queryResolver) Bab(ctx context.Context, babID int, email *string) (*models.Bab, error) {
	var bab models.Bab
	result := r.DB.Preload("SubBabs", util.FuncOrder).
		Preload("SubBabs.Contents", util.FuncOrder).
		Preload("SubBabs.Contents.Logs", "Email = ?", email).
		Find(&bab, babID)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &bab, nil
}

func (r *queryResolver) Material(ctx context.Context, contentID int, email *string) (*models.Material, error) {
	var material models.Material
	result := r.DB.Preload("Content.Logs", "Email = ?", email).Find(&material, contentID)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	return &material, nil
}

func (r *queryResolver) Test(ctx context.Context, contentID int, email *string) (*models.Test, error) {
	var test models.Test
	result := r.DB.Preload("Content.Logs", "Email = ?", email).Preload("Questions", util.FuncOrder).Preload("Logs", "Email = ?", email).Find(&test, contentID)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	return &test, nil
}

func (r *queryResolver) IsCheck(ctx context.Context, contentID int, email string) (bool, error) {
	var log models.ContentLog
	result := r.DB.Where(&models.ContentLog{Email: email, ContentID: contentID}).Find(&log)

	if result.Error != nil {
		return false, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *queryResolver) SubBabs(ctx context.Context) ([]*models.SubBab, error) {
	var subBabs []*models.SubBab
	result := r.DB.Preload("Bab").Find(&subBabs)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return subBabs, nil
}

func (r *testResolver) Log(ctx context.Context, obj *models.Test) (*models.ExamLog, error) {
	if len(obj.Logs) > 0 {
		return &obj.Logs[0], nil
	}

	return nil, nil
}

// Content returns generated.ContentResolver implementation.
func (r *Resolver) Content() generated.ContentResolver { return &contentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Test returns generated.TestResolver implementation.
func (r *Resolver) Test() generated.TestResolver { return &testResolver{r} }

type contentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type testResolver struct{ *Resolver }
