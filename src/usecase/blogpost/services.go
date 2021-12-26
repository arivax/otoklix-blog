package blogpost

type BlogPostSvc struct{}

var Svc = BlogPostSvc{}

func (o *BlogPostSvc) FindAll() ([]BlogPostRes, error) {
	res := []BlogPostRes{}

	rows, err := Repo.FindAll()
	if err != nil {
		return res, err
	}
	for _, row := range rows {
		res = append(res, BlogPostRes(row))
	}
	return res, nil
}

func (o *BlogPostSvc) FindById(id string) (BlogPostRes, error) {
	res := BlogPostRes{}

	row, err := Repo.FindById(id)
	if err != nil {
		return res, err
	}
	res = BlogPostRes(row)
	return res, nil
}

func (o *BlogPostSvc) CreateOne(data *BlogPostReq) (BlogPostRes, error) {
	res := BlogPostRes{}

	row, err := Repo.CreateOne(data)
	if err != nil {
		return res, err
	}
	res = BlogPostRes(row)
	return res, nil
}

func (o *BlogPostSvc) UpdateOne(data *BlogPostReq, id string) (BlogPostRes, error) {
	res := BlogPostRes{}

	row, err := Repo.UpdateOne(data, id)
	if err != nil {
		return res, err
	}
	res = BlogPostRes(row)
	return res, nil
}

func (o *BlogPostSvc) DeleteOne(data *BlogPostReq, id string) (BlogPostRes, error) {
	res := BlogPostRes{}

	row, err := Repo.DeleteOne(id)
	if err != nil {
		return res, err
	}
	res = BlogPostRes(row)
	return res, nil
}
