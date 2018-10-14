package ice

import (
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	LogFile            string
	CodeDir            string
	RepositoryProvider string
	ApiToken           string
}

const CONFIG_FILE_NAME = "config.json"
const ICE_FOLDER = ".ice"

var appConfig *Configuration

func InitializeAppConfigs() {
	// don't re-initialize the configs
	if appConfig != nil {
		return
	}
	// ensure th config file is created
	appConfig = newConfiguration()
	createDefaultConfigFile(appConfig)

	err := appConfig.readFromFile(GetConfigFileLocation())

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetConfig() *Configuration {
	return appConfig
}

func getUserHomeDir() string {
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return homeDir
}

func GetConfigFileLocation() string {
	return filepath.Join(getUserHomeDir(), ICE_FOLDER, CONFIG_FILE_NAME)
}

func newConfiguration() *Configuration {
	homeDir := getUserHomeDir()

	return &Configuration{
		LogFile:            filepath.Join(homeDir, ICE_FOLDER, "logger"),
		CodeDir:            filepath.Join(homeDir, ICE_FOLDER, "code"),
		RepositoryProvider: "github4",
	}
}

func (c *Configuration) readFromFile(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, c)
}

func (c *Configuration) writeToFile(fileName string) error {
	b, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, b, 0755)
}

// createDefaultConfigFile creates a new
func createDefaultConfigFile(defaults *Configuration) {
	fileName := GetConfigFileLocation()
	// return if the file exists
	_, err := os.Stat(fileName)
	if err == nil {
		return
	}

	path := filepath.Dir(fileName)
	_, err = os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	} else if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err.Error())
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := defaults.writeToFile(fileName); err != nil {
		log.Fatalln(err.Error())
	}

	file.Close()
}
