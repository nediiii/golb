package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/middlewares"
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

func (r *mutationResolver) Auth(ctx context.Context, username *string, password *string, token *string) (string, error) {
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
					newClaims := utils.GenerateClaims(v.Slug, v.ID, v.RoleID)
					tokenStr, err := utils.GenerateTokenWithClaims(newClaims)
					if err == nil {
						return tokenStr, nil
					}
				}
			}
		}
		return "", errors.New("Authentication fail, please retry with correct token")
	} else if username != nil && password != nil {
		// use username and password
		tx := services.DB
		v := &models.User{}
		if !tx.Where("name = ?", username).First(v).RecordNotFound() {
			if utils.IsCorrect(v.Password, *password) {
				claims := utils.GenerateClaims(v.Slug, v.ID, v.RoleID)
				tokenStr, err := utils.GenerateTokenWithClaims(claims)
				if err == nil {
					return tokenStr, nil
				}
			}
		}
		return "", errors.New("Authentication fail, please retry with correct username and password")
	}
	return "", errors.New("Invalid input! You must send a token or (username and password)")
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

func (r *mutationResolver) CreateUser(ctx context.Context, slug string, password string, name string, role string, email *string, bio *string) (*models.User, error) {
	if err := utils.CheckPermission(ctx, models.RoleEditor); err != nil {
		return nil, err
	}

	obj := &models.User{}

	obj.Slug = slug
	obj.Name = name
	obj.Password = password
	obj.RoleID = utils.String2Uint(role)
	if email != nil {
		obj.Email = *email
	}
	if bio != nil {
		obj.Bio = *bio
	}

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

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, slug *string, name *string, email *string, oldPassword *string, newPassword *string, bio *string) (*models.User, error) {
	obj := &models.User{}

	tx := services.DB
	if tx.First(obj, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}

	if oldPassword != nil || newPassword != nil {
		if oldPassword == nil || newPassword == nil {
			return nil, errors.New("`oldPassword`和`newPassword`必须同时存在或者同时不存在")
		}
		if len(*newPassword) < 8 || len(*newPassword) > 64 {
			return nil, errors.New("新密码长度不合规,请确认后重试")
		}
		if !utils.IsCorrect(obj.Password, *oldPassword) {
			return nil, errors.New("旧密码错误, 请确认后重试")
		}
		obj.Password = *newPassword
	}

	if slug != nil {
		obj.Slug = *slug
	}
	if email != nil {
		obj.Email = *email
	}
	if name != nil {
		obj.Name = *name
	}
	if bio != nil {
		obj.Bio = *bio
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

func (r *mutationResolver) CreatePost(ctx context.Context, authors []string, commentable *bool, excerpt *string, featured *bool, html string, image *string, metaTitle *string, metaDescription *string, markdown string, primaryAuthorID string, publishedBy *string, publishedAt *string, paged *bool, slug string, status *string, tags []string, title string) (*models.Post, error) {
	obj := &models.Post{}

	obj.Slug = slug
	obj.Title = title
	obj.Markdown = markdown
	obj.HTML = html
	obj.PrimaryAuthorID = utils.String2Uint(primaryAuthorID)
	services.DB.Where("id IN (?)", tags).Find(&obj.Tags)
	services.DB.Where("id IN (?)", authors).Find(&obj.Authors)
	services.DB.Find(&obj.PrimaryAuthor, primaryAuthorID)

	if excerpt != nil && obj.Excerpt != *excerpt {
		obj.Excerpt = *excerpt
	}
	if featured != nil && obj.Featured != *featured {
		obj.Featured = *featured
	}
	if paged != nil && obj.Paged != *paged {
		obj.Paged = *paged
	}
	if commentable != nil && obj.Commentable != *commentable {
		obj.Commentable = *commentable
	}
	if image != nil && obj.Image != *image {
		obj.Image = *image
	}
	if metaTitle != nil && obj.MetaTitle != *metaTitle {
		obj.MetaTitle = *metaTitle
	}
	if metaDescription != nil && obj.MetaDescription != *metaDescription {
		obj.MetaDescription = *metaDescription
	}
	if status != nil && obj.Status != *status {
		obj.Status = *status
	}
	if publishedBy != nil && obj.PublishedBy != utils.String2Uint(*publishedBy) {
		obj.PublishedBy = utils.String2Uint(*publishedBy)
	}
	if publishedAt != nil && obj.PublishedAt != utils.UnixString2Time(*publishedAt) {
		obj.PublishedAt = utils.UnixString2Time(*publishedAt)
	}

	services.DB.Where("id IN (?)", tags).Find(&obj.Tags)
	services.DB.Model(obj).Association("Tags").Append(obj.Tags)
	services.DB.Where("id IN (?)", authors).Find(&obj.Authors)
	services.DB.Model(obj).Association("Authors").Append(obj.Authors)

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

func (r *mutationResolver) UpdatePost(ctx context.Context, authors []string, commentable *bool, excerpt *string, featured *bool, html *string, id string, image *string, markdown *string, metaTitle *string, metaDescription *string, paged *bool, primaryAuthorID *string, publishedBy *string, publishedAt *string, slug *string, status *string, tags []string, title *string) (*models.Post, error) {
	obj := &models.Post{}

	tx := services.DB
	if tx.First(obj, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}

	if slug != nil && obj.Slug != *slug {
		obj.Slug = *slug
	}
	if title != nil && obj.Title != *title {
		obj.Title = *title
	}
	if markdown != nil && obj.Markdown != *markdown {
		obj.Markdown = *markdown
	}
	if html != nil && obj.HTML != *html {
		obj.HTML = *html
	}
	if primaryAuthorID != nil && obj.PrimaryAuthorID != utils.String2Uint(*primaryAuthorID) {
		obj.PrimaryAuthorID = utils.String2Uint(*primaryAuthorID)
	}
	if excerpt != nil && obj.Excerpt != *excerpt {
		obj.Excerpt = *excerpt
	}
	if featured != nil && obj.Featured != *featured {
		obj.Featured = *featured
	}
	if paged != nil && obj.Paged != *paged {
		obj.Paged = *paged
	}
	if commentable != nil && obj.Commentable != *commentable {
		obj.Commentable = *commentable
	}
	if image != nil && obj.Image != *image {
		obj.Image = *image
	}
	if metaTitle != nil && obj.MetaTitle != *metaTitle {
		obj.MetaTitle = *metaTitle
	}
	if metaDescription != nil && obj.MetaDescription != *metaDescription {
		obj.MetaDescription = *metaDescription
	}
	if status != nil && obj.Status != *status {
		obj.Status = *status
	}
	if publishedBy != nil && obj.PublishedBy != utils.String2Uint(*publishedBy) {
		obj.PublishedBy = utils.String2Uint(*publishedBy)
	}
	if publishedAt != nil && obj.PublishedAt != utils.UnixString2Time(*publishedAt) {
		obj.PublishedAt = utils.UnixString2Time(*publishedAt)
	}

	services.DB.Where("id IN (?)", tags).Find(&obj.Tags)
	middlewares.GetDataloaderFromContext(ctx).PostTagsLoader.Clear(obj.ID) // clean the dataloader cache, makesure data is newest.
	services.DB.Model(obj).Association("Tags").Replace(obj.Tags)
	services.DB.Where("id IN (?)", authors).Find(&obj.Authors)
	services.DB.Model(obj).Association("Authors").Replace(obj.Authors)

	var err gorm.Errors
	if tx.Save(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return obj, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, nickname string, email string, target string, content string, postID string, parentID string) (*models.Comment, error) {
	tx := services.DB
	obj := &models.Comment{}

	obj.Nickname = nickname
	obj.Email = email
	obj.Target = target
	obj.Content = content
	obj.PostID = utils.String2Uint(postID)
	obj.ParentID = utils.String2Uint(parentID)

	ginContext := middlewares.GetGinContextFromContext(ctx)
	obj.IP = ginContext.ClientIP()
	obj.Agent = ginContext.Request.Header.Get("User-Agent")

	var err gorm.Errors
	if err = tx.Create(obj).GetErrors(); len(err) > 0 {
		return nil, err
	}

	// emial notice
	// @guest reply your comment. plz visit @site to check it.
	targetPost := &models.Post{}
	tx.Model(targetPost).Find(&targetPost, postID)
	targetComment := &models.Comment{}
	var replyRecipient string
	if tx.Model(targetComment).Find(&targetComment, parentID).RecordNotFound() {
		// notice the post owner
		var u models.User
		tx.Model(targetPost).Related(&u, "PrimaryAuthorID")
		replyRecipient = u.Email
	} else {
		// notice the comment owner
		replyRecipient = targetComment.Email
	}
	go services.Reply(email, replyRecipient, targetPost.Slug)
	return obj, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
