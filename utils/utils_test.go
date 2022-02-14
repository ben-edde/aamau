package utils

import (
	"os"
	"reflect"
	"testing"
)

func Test_read_config(t *testing.T) {
	config_path := "../cfg/test_config.yaml"
	expected_result := map[string]string{
		"USERNAME": "1",
		"PASSWORD": "2",
		"SERVER":   "3",
		"PORT":     "4",
		"DATABASE": "5"}

	executed_result := Read_config(&config_path)

	if !reflect.DeepEqual(expected_result, executed_result) {
		t.Error("read_config() failed.")
	}
}

func Test_check_path_exists(t *testing.T) {
	err := os.Mkdir("testing", 0754)
	if err != nil {
		t.Error("mkdir failed.")
	}
	path_existed, err := Check_path_exists("testing/")
	if path_existed != true || err != nil {
		t.Error("check_path_exists failed.")
	}
	err = os.Remove("testing")
	if err != nil {
		t.Error("rmdir failed.")
	}
}

func Test_check_path_exists_dir_not_exist(t *testing.T) {
	path_existed, err := Check_path_exists("testing/")
	if path_existed != false || err != nil {
		t.Error("check_path_exists failed.")
	}
}
