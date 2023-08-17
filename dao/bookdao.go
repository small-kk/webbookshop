package dao

import (
	"book_mall/model"
	"book_mall/repository"
	"strconv"
)

// GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	books := []*model.Book{}
	err := repository.DB.Find(&books).Error
	return books, err
}

// AddBook 向数据库中添加一本图书
func AddBook(book *model.Book) error {
	err := repository.DB.Create(&book).Error
	return err
}

// DeleteBook 从数据库中删除一本图书
func DeleteBook(book *model.Book) error {
	err := repository.DB.Delete(book).Error
	return err
}

// GetBookByID 根据bookID获取一本图书信息
func GetBookByID(bookID int) (*model.Book, error) {
	book := &model.Book{}
	err := repository.DB.Where("id=?", bookID).Find(&book).Error
	return book, err
}

// UpdatesBook 更新图书
func UpdatesBook(book *model.Book) error {
	err := repository.DB.Updates(book).Error
	return err
}

// GetPageBooks 获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	//获取数据库中图书总数
	var count int64
	err := repository.DB.Model(model.Book{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	//每页显示四条记录
	pageSize := 4
	//获取总页数
	var totalPageNo int
	if (int(count) % pageSize) == 0 {
		totalPageNo = int(count) / pageSize
	} else {
		totalPageNo = int(count)/pageSize + 1
	}

	ipageNo, _ := strconv.Atoi(pageNo)

	//获取每一页图书信息
	books := []*model.Book{}
	err = repository.DB.Limit(pageSize).Offset((ipageNo - 1) * pageSize).Find(&books).Error

	if err != nil {
		return nil, err
	}

	page := &model.Page{
		Books:       books,
		PageNo:      ipageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: int(count),
	}
	return page, nil
}

func GetPageBooksByPrice(pageNo, minPrice, maxPrice string) (*model.Page, error) {
	//查询在minPrice 和 maxPrice的价格之间的图书数量
	var count int64
	err := repository.DB.Model(model.Book{}).Where("price between ? and ? ", minPrice, maxPrice).Count(&count).Error
	if err != nil {
		return nil, err
	}
	//每页显示四条记录
	pageSize := 4
	//获取在价格区间的图书总页数
	var totalPageNo int
	if (int(count) % pageSize) == 0 {
		totalPageNo = int(count) / pageSize
	} else {
		totalPageNo = int(count)/pageSize + 1
	}
	ipageNo, _ := strconv.Atoi(pageNo)
	//获取每一页图书信息
	books := []*model.Book{}
	err = repository.DB.Where("price between ? and ? ", minPrice, maxPrice).Limit(pageSize).Offset((ipageNo - 1) * pageSize).Find(&books).Error

	if err != nil {
		return nil, err
	}

	page := &model.Page{
		Books:       books,
		PageNo:      ipageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: int(count),
	}
	return page, nil
}
