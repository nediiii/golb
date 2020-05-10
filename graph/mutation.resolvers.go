package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"golb/services"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func (r *mutationResolver) Auth(ctx context.Context, username *string, password *string, token *string) (*model.Jwt, error) {
	var jwtKey = []byte("golb.sys.jwt.key")

	if token != nil {
		parseToken, err := jwt.Parse(*token, func(tk *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", tk.Header["alg"])
			}

			// jwtKey is a []byte containing your secret, e.g. []byte("my_secret_key")
			return jwtKey, nil
		})
		if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
			fmt.Println(claims["exp"], claims["iss"])
			fmt.Println("valid: ", parseToken.Valid)
			// aud exp jti iat iss nbf sub
			return nil, nil
		}
		return nil, err
	} else if username != nil && password != nil {
		claims := &jwt.StandardClaims{
			Audience:  *username,                               // 受众
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // 失效时间
			Id:        "1",                                     // 编号
			IssuedAt:  time.Now().Unix(),                       // 签发时间
			Issuer:    "golb.sys",                              // 签发人
			NotBefore: time.Now().Unix(),                       // 生效时间
			Subject:   "login",                                 // 主题
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		str, err := token.SignedString(jwtKey)
		if err != nil {
			fmt.Print(err.Error())
		}
		return &model.Jwt{ExpireAt: strconv.FormatInt(claims.ExpiresAt, 10), Token: str}, nil
	}
	return nil, errors.New("Invalid input! You must send a token or (username and password)")
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	// ioutil.WriteFile("statics/"+file.Filename, content, 0644)
	return &model.File{
		ID:      1,
		Name:    file.Filename,
		Content: string(content),
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
	s := &models.Setting{}
	s.Key = key
	s.Value = value
	var err gorm.Errors
	err = services.DB.Create(s).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return s, nil
}

func (r *mutationResolver) DeleteSetting(ctx context.Context, id string) (bool, error) {
	var err gorm.Errors
	err = services.DB.Where("id = ?", id).Delete(models.Setting{}).GetErrors()
	if len(err) > 0 {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateSetting(ctx context.Context, id string, key string, value string) (*models.Setting, error) {
	v := &models.Setting{}
	tx := services.DB
	if tx.First(v, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}
	v.Key = key
	v.Value = value
	var err gorm.Errors
	tx.Save(v).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return v, nil
}

func (r *mutationResolver) CreateRole(ctx context.Context, name string, description *string) (*models.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRole(ctx context.Context, id string, name *string, description *string) (*models.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, slug string, name string, password string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, slug *string, name *string, password *string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTag(ctx context.Context, slug string, name string, description *string) (*models.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTag(ctx context.Context, id string, slug *string, name *string, description *string) (*models.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, slug string, title string, markdown string, html string, primaryAuthorID string, tags []string, authors []string, excerpt *string, fetured *bool, paged *bool, publishedBy *string, image *string, language *string, status *string) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, slug *string, title *string, markdown *string, html *string, primaryAuthorID *string, tags []string, authors []string, excerpt *string, fetured *bool, paged *bool, publishedBy *string, image *string, language *string, status *string) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
