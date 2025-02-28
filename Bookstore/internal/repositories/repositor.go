package repositories

import (
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/models"
	"gorm.io/gorm"  // GORM baza dana bilen islemek ucin biblioteka. 
)                   // Ctoby db dannylar bilen sqlite girman su yerde islap bolar yaly

type BookRepository interface {          //interfeys gerekmi su yerde. Service-em su zatlary etyana menzeya
	GetAllBooks() ([]models.Book, error)  // database hemme kitaplary alya
	GetBookByID(id uint) (*models.Book, error)  // database kitapyn id'sine gora alya
	CreateBook(book *models.Book) error    //database taze kitap gosyar. Id boyunca
	UpdateBook(id uint, book *models.Book) error // database kitapyn id'sine gora kitaby obnowit etya
	DeleteBook(id uint) error // id boyunca kitaby ayyryar
}

type bookRepository struct {   //database catmak bookRepisitory
	db *gorm.DB 
} 

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}   // name gaytaryanyny dusunup bilmedim
}

func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book  // kitaplar massiw gornusde. Book structuradaky tipleri bilen (ady, awtory ...)
	err := r.db.Find(&books).Error // Find baza dannydan Book strukturaky kitaplary baza dannydan alyp books a dakyar 
	return books, err    // Sorag? baza dannydan alyp books a dakyas. Baza danna maglumatlay hacan gizirdik? Ondan biz alar yaly?
}

func (r *bookRepository) GetBookByID(id uint) (*models.Book, error) {   
	var book models.Book   // Boook strukturady tiply kitaplary book peremenna dakyar
	err := r.db.First(&book, id).Error  // Firs id boyunca Book struct daky kitaby  baza dannydan alyar. Klient tarapyndan soralyan id boyunca kitap alynyar
		return &book, err   //Sorag? Aydaly= spisok 1 den 10 a cenli kitap. 5 nji kitaby saylady klient. 5 nji kitaby alyp bermeli kod. Sonda id boyunca zapros bolyamy su? 
}                           // Beyle bolyan bolsa onda  Book strukturadaky ahli tipler boyunca zapros bolyada klien bir kitap saylasa. 

func (r *bookRepository) CreateBook(book *models.Book) error {
  	return r.db.Create(book).Error   // database taze kitap gosup bazadan alyp  ony book peremmenna dakyar.
}                                    // Umumy sorag baza danna kitaplary nirde yukledik biz. Gorup bilemok mena. Yukleyan koda yiok yalak barde?
                                    // Bazadan alyp otyrys yone. Baza kim gosdy? Nadip gosdy?

func (r *bookRepository) UpdateBook(id uint, book *models.Book) error {   //UpdateBook su yerde islemek ucin da.
	return r.db.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error // Gormdaky update bolsa bize databasada obnowit etmek ucin kitaby.
	                                                            /* Model kitap strukturasynda kitapyn id'sine gora kitaby alyp berya. Where id berlen kitabyn
	id sini denesdiryar  Deng bolsa update obnowit edip r *bookRepositora gaytaryarmy. Deng dal bolsa innika  gecyami yada error beryami*/
}

func (r *bookRepository) DeleteBook(id uint) error {  // Aslynda dusunip otyryn yone dlya utocneniya. Acylan baza deletebookyng tipynda bolya.Name ucin? Cunlasdyryp dusundiray! Mozhet men yalnysh dusunup otyrandyryn
	return r.db.Delete(&models.Book{}, id).Error   // Databasada id boyunca kitaby udalit edip gaytaryar. Error ya sol baza danny bilen catylyp bilmese beryarmi.
}

// Asakda bazadannyny  dokcera gosjak boldum. Ce ta bolman galda. Konteyner acyp bomady oytyan dockerda sunda. Osipka berip galyar. Interneding yiok diyya
/*
func initDB() {
	var err error
	dsn := "host=localhost user=sqlite password=123 dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Baku"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic(err)
	}
	db.AutoMigrate(&models.Book{})
}
*/