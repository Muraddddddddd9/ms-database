package s3

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/Muraddddddddd9/ms-database/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	keyPrefix = "%s/"
	bucket    = "diplomas"
)

type S3Server struct {
	s3 *s3.S3
}

func Connect() (*S3Server, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	awsConfig := &aws.Config{
		Endpoint:         aws.String(cfg.S3_ENDPOINT),
		Region:           aws.String(cfg.S3_REGION),
		Credentials:      credentials.NewStaticCredentials(cfg.S3_CREDENTIALS_ID, cfg.S3_CREDENTIALS_SECRET, cfg.S3_CREDENTIALS_TOKEN),
		DisableSSL:       aws.Bool(cfg.S3_DISABLE_SSL),
		S3ForcePathStyle: aws.Bool(cfg.S3_FORCE_PATH_STYLE),
	}

	s3Sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	serv := &S3Server{s3: s3.New(s3Sess)}

	_, err = serv.s3.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		err = serv.CreateBucket()
		if err != nil {
			return nil, err
		}
	}

	return serv, nil
}

type PdfInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func (s *S3Server) CreateBucket() error {
	_, err := s.s3.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return fmt.Errorf("Failed to create bucket: %v", err)
	}

	return nil
}

func (s *S3Server) ListFileS3(id string) (*[]PdfInfo, error) {
	result, err := s.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(fmt.Sprintf(keyPrefix, id)),
	})

	if err != nil {
		return nil, err
	}

	var pdfs []PdfInfo
	for _, obj := range result.Contents {
		pdfs = append(pdfs, PdfInfo{
			Name: *obj.Key,
			Size: *obj.Size,
		})
	}

	return &pdfs, nil
}

func (s *S3Server) GetFileS3(filename, id string) (*s3.GetObjectOutput, error) {
	result, err := s.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf(keyPrefix, id) + filename),
	})

	if err != nil {
		return nil, err
	}

	return result, err
}

func (s *S3Server) UploadFileS3(file multipart.File, header *multipart.FileHeader, id string) error {
	filename := transliterate(header.Filename)

	_, err := s.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf(keyPrefix, id) + filename),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}

func transliterate(input string) string {
	rules := map[rune]string{
		'А': "A", 'а': "a",
		'Б': "B", 'б': "b",
		'В': "V", 'в': "v",
		'Г': "G", 'г': "g",
		'Д': "D", 'д': "d",
		'Е': "E", 'е': "e",
		'Ё': "Yo", 'ё': "yo",
		'Ж': "Zh", 'ж': "zh",
		'З': "Z", 'з': "z",
		'И': "I", 'и': "i",
		'Й': "Y", 'й': "y",
		'К': "K", 'к': "k",
		'Л': "L", 'л': "l",
		'М': "M", 'м': "m",
		'Н': "N", 'н': "n",
		'О': "O", 'о': "o",
		'П': "P", 'п': "p",
		'Р': "R", 'р': "r",
		'С': "S", 'с': "s",
		'Т': "T", 'т': "t",
		'У': "U", 'у': "u",
		'Ф': "F", 'ф': "f",
		'Х': "Kh", 'х': "kh",
		'Ц': "Ts", 'ц': "ts",
		'Ч': "Ch", 'ч': "ch",
		'Ш': "Sh", 'ш': "sh",
		'Щ': "Shch", 'щ': "shch",
		'Ъ': "", 'ъ': "",
		'Ы': "Y", 'ы': "y",
		'Ь': "", 'ь': "",
		'Э': "E", 'э': "e",
		'Ю': "Yu", 'ю': "yu",
		'Я': "Ya", 'я': "ya",
	}

	var result strings.Builder
	for _, char := range input {
		if val, ok := rules[char]; ok {
			result.WriteString(val)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
