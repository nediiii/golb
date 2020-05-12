package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"golb/services"
	"golb/utils"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func (r *mutationResolver) Auth(ctx context.Context, username *string, password *string, token *string) (*model.Jwt, error) {
	if token != nil {
		// use token
		parseToken, err := utils.JwtParse(*token)
		if err == nil {
			// 生成新token
			if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
				tx := services.DB
				v := &models.User{}
				name := claims["aud"].(string)
				if !tx.Where("name = ?", name).First(v).RecordNotFound() {
					newClaims := utils.GenerateClaims(v.Slug, v.ID)
					tokenStr, err := utils.GenerateTokenWithClaims(newClaims)
					if err == nil {
						jwt := &model.Jwt{ExpireAt: strconv.FormatInt(newClaims.ExpiresAt, 10), Token: tokenStr}
						return jwt, nil
					}
				}
			}
		}
		return nil, errors.New("Authentication fail, please retry with correct token")
	} else if username != nil && password != nil {
		// use username and password
		tx := services.DB
		v := &models.User{}
		if !tx.Where("name = ?", username).First(v).RecordNotFound() {
			if utils.IsCorrect(v.Password, *password) {
				claims := utils.GenerateClaims(v.Slug, v.ID)
				tokenStr, err := utils.GenerateTokenWithClaims(claims)
				if err == nil {
					jwt := &model.Jwt{ExpireAt: strconv.FormatInt(claims.ExpiresAt, 10), Token: tokenStr}
					return jwt, nil
				}
			}
		}
		return nil, errors.New("Authentication fail, please retry with correct username and password")
	}
	return nil, errors.New("Invalid input! You must send a token or (username and password)")
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	h := sha256.New()
	h.Write(content)
	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	ioutil.WriteFile("statics/"+timeStr+".png", content, 0644)

	// ioutil.WriteFile("statics/"+file.Filename, content, 0644)
	return &model.File{
		ID:      1,
		Name:    file.Filename,
		Content: string(content),
		Hash:    hex.EncodeToString(h.Sum(nil)),
		URL:     "http://localhost:8090/statics/" + timeStr + ".png",
	}, nil
}

func (r *mutationResolver) SingleUploadWithPayload(ctx context.Context, req model.UploadFile) (*model.File, error) {
	content, err := ioutil.ReadAll(req.File.File)
	if err != nil {
		return nil, err
	}
	return &model.File{
		ID:      1,
		Name:    req.File.Filename,
		Content: string(content),
	}, nil
}

func (r *mutationResolver) MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	if len(files) == 0 {
		return nil, errors.New("empty list")
	}
	var resp []*model.File
	for i := range files {
		content, err := ioutil.ReadAll(files[i].File)
		if err != nil {
			return []*model.File{}, err
		}
		resp = append(resp, &model.File{
			ID:      i + 1,
			Name:    files[i].Filename,
			Content: string(content),
		})
	}
	return resp, nil
}

func (r *mutationResolver) MultipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
	if len(req) == 0 {
		return nil, errors.New("empty list")
	}
	var resp []*model.File
	for i := range req {
		content, err := ioutil.ReadAll(req[i].File.File)
		if err != nil {
			return []*model.File{}, err
		}
		resp = append(resp, &model.File{
			ID:      i + 1,
			Name:    req[i].File.Filename,
			Content: string(content),
		})
	}
	return resp, nil
}

func (r *mutationResolver) CreateSetting(ctx context.Context, key string, value string) (*models.Setting, error) {
	obj := &models.Setting{}

	obj.Key = key
	obj.Value = value

	var err gorm.Errors
	if err = services.DB.Create(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) DeleteSetting(ctx context.Context, id string) (bool, error) {
	obj := &models.Setting{}

	var err gorm.Errors
	if err = services.DB.Where("id = ?", id).Delete(obj).GetErrors(); len(err) > 0 {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateSetting(ctx context.Context, id string, key *string, value *string) (*models.Setting, error) {
	obj := &models.Setting{}

	tx := services.DB
	if tx.First(obj, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}

	if key != nil {
		obj.Key = *key
	}
	if value != nil {
		obj.Value = *value
	}

	var err gorm.Errors
	if tx.Save(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) CreateRole(ctx context.Context, name string, description *string) (*models.Role, error) {
	obj := &models.Role{}

	obj.Name = name
	if description != nil {
		obj.Description = *description
	}

	var err gorm.Errors
	if err = services.DB.Create(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (bool, error) {
	obj := &models.Role{}

	var err gorm.Errors
	if err = services.DB.Where("id = ?", id).Delete(obj).GetErrors(); len(err) > 0 {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateRole(ctx context.Context, id string, name *string, description *string) (*models.Role, error) {
	obj := &models.Role{}

	tx := services.DB
	if tx.First(obj, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}

	if name != nil {
		obj.Name = *name
	}
	if description != nil {
		obj.Description = *description
	}

	var err gorm.Errors
	if tx.Save(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, slug string, name string, password string) (*models.User, error) {
	obj := &models.User{}

	obj.Slug = slug
	obj.Name = name
	obj.Password = password

	var err gorm.Errors
	if err = services.DB.Create(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	obj := &models.User{}

	var err gorm.Errors
	if err = services.DB.Where("id = ?", id).Delete(obj).GetErrors(); len(err) > 0 {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, slug *string, name *string, password *string) (*models.User, error) {
	obj := &models.User{}

	tx := services.DB
	if tx.First(obj, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}

	if slug != nil {
		obj.Slug = *slug
	}
	if name != nil {
		obj.Name = *name
	}
	if password != nil {
		obj.Password = *password
	}

	var err gorm.Errors
	if tx.Save(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) CreateTag(ctx context.Context, slug string, name string, description *string) (*models.Tag, error) {
	obj := &models.Tag{}

	obj.Slug = slug
	obj.Name = name
	if description != nil {
		obj.Description = *description
	}

	var err gorm.Errors
	if err = services.DB.Create(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	obj := &models.Tag{}

	var err gorm.Errors
	if err = services.DB.Where("id = ?", id).Delete(obj).GetErrors(); len(err) > 0 {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateTag(ctx context.Context, id string, slug *string, name *string, description *string) (*models.Tag, error) {
	obj := &models.Tag{}

	tx := services.DB
	if tx.First(obj, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}

	if slug != nil {
		obj.Slug = *slug
	}
	if name != nil {
		obj.Name = *name
	}
	if description != nil {
		obj.Description = *description
	}

	var err gorm.Errors
	if tx.Save(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, slug string, title string, markdown string, html string, primaryAuthorID string, tags []string, authors []string, excerpt *string, fetured *bool, paged *bool, publishedBy *string, image *string, language *string, status *string) (*models.Post, error) {
	obj := &models.Post{}

	obj.Slug = slug
	obj.Title = title
	obj.Markdown = markdown
	obj.HTML = html

	var err gorm.Errors
	if err = services.DB.Create(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	obj := &models.Post{}

	var err gorm.Errors
	if err = services.DB.Where("id = ?", id).Delete(obj).GetErrors(); len(err) > 0 {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, slug *string, title *string, markdown *string, html *string, primaryAuthorID *string, tags []string, authors []string, excerpt *string, fetured *bool, paged *bool, publishedBy *string, image *string, language *string, status *string) (*models.Post, error) {
	obj := &models.Post{}

	tx := services.DB
	if tx.First(obj, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}

	if slug != nil {
		obj.Slug = *slug
	}
	if title != nil {
		obj.Title = *title
	}
	if html != nil {
		obj.HTML = *html
	}
	if markdown != nil {
		obj.Markdown = *markdown
	}

	var err gorm.Errors
	if tx.Save(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
