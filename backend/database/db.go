package database

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/DylanCoon99/portfolio/backend/models"
	//"fmt"
)

var DB *gorm.DB



func ConnectDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
		return err
	}

	

	// DROP TABLES
	
	//DB.Migrator().DropTable(&models.Project{})

	
	


	
	err = DB.AutoMigrate(&models.Project{})    // check
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	

	log.Println("Database connected and migrated successfully!")
	return nil
}



func GetDB() *gorm.DB {
	return DB
}



func GetAllProjects() ([]models.Project, error) {

	var projects []models.Project

	if err := DB.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
} 



func GetProjectByID(project_id string) (*models.Project, error){

	var project models.Project

	if err := DB.Find(&project, project_id).Error; err != nil {
		return nil, err
	}

	return &project, nil
}



func CreateProject(project *models.Project) error {
	result := DB.Create(project)
	return result.Error
}



func UploadImage(id, path string) error {

	result := DB.Model(&models.Project{}).Where("project_id = ?", id).Update("image_url", path)
	return result.Error
}


func GetImagePath(project_id string) (string, error) {

	var imagePath string


	if err := DB.Table("projects").Select("image_url").Where("project_id = ?", project_id).Scan(&imagePath).Error; err != nil {
		return "", err
	}

	return imagePath, nil
}