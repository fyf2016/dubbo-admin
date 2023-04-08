/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/config_center"
	"dubbo.apache.org/dubbo-go/v3/metadata/report"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/apache/dubbo-admin/pkg/admin/constant"
	_ "github.com/apache/dubbo-admin/pkg/admin/imports"
	"github.com/apache/dubbo-admin/pkg/admin/model"
	"gopkg.in/yaml.v2"
)

const conf = "./conf/dubboadmin.yml"

type Config struct {
	Admin      Admin      `yaml:"admin"`
	Prometheus Prometheus `yaml:"prometheus"`
}

type Prometheus struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type Admin struct {
	ConfigCenter   string        `yaml:"config-center"`
	MetadataReport AddressConfig `yaml:"metadata-report"`
	Registry       AddressConfig `yaml:"registry"`
	MysqlDsn       string        `yaml:"mysql-dsn"`
}

var (
	ConfigCenter         config_center.DynamicConfiguration
	RegistryCenter       registry.Registry
	MetadataReportCenter report.MetadataReport

	DataBase *gorm.DB // for service mock

	Group string
)

var (
	PrometheusIp   string
	PrometheusPort string
)

func LoadConfig() {
	path, err := filepath.Abs(conf)
	if err != nil {
		path = filepath.Clean(conf)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var config Config
	yaml.Unmarshal(content, &config)

	address := config.Admin.ConfigCenter
	registryAddress := config.Admin.Registry.Address
	metadataReportAddress := config.Admin.MetadataReport.Address

	loadDatabaseConfig(config.Admin.MysqlDsn)

	PrometheusIp = config.Prometheus.Ip
	PrometheusPort = config.Prometheus.Port
	if PrometheusIp == "" {
		PrometheusIp = "127.0.0.1"
	}
	if PrometheusPort == "" {
		PrometheusPort = "9090"
	}

	if len(address) > 0 {
		c := newAddressConfig(address)
		factory, err := extension.GetConfigCenterFactory(c.getProtocol())
		if err != nil {
			panic(err)
		}
		url, err := c.toURL()
		if err != nil {
			panic(err)
		}
		ConfigCenter, err = factory.GetDynamicConfiguration(url)
		Group = url.GetParam(constant.GroupKey, constant.DefaultGroup)
		if err != nil {
			log.Print("No configuration found in config center.")
		}
		properties, err := ConfigCenter.GetProperties(constant.DubboPropertyKey)
		if err != nil {
			log.Print("No configuration found in config center.")
		}
		if len(properties) > 0 {
			for _, property := range strings.Split(properties, "\n") {
				if strings.HasPrefix(property, constant.RegistryAddressKey) {
					registryAddress = strings.Split(property, "=")[1]
				}
				if strings.HasPrefix(property, constant.MetadataReportAddressKey) {
					metadataReportAddress = strings.Split(property, "=")[1]
				}
			}
		}
	}
	if ConfigCenter == nil {
		if len(registryAddress) == 0 {
			panic("registry address can not be empty")
		}
		c := newAddressConfig(registryAddress)
		url, err := c.toURL()
		if err != nil {
			panic(err)
		}
		factory, err := extension.GetConfigCenterFactory(c.getProtocol())
		if err != nil {
			log.Print("No configuration found in config center.")
		}
		ConfigCenter, err = factory.GetDynamicConfiguration(url)
		if err != nil {
			panic(err)
		}
		Group = url.GetParam(constant.GroupKey, constant.DefaultGroup)
	}
	if len(registryAddress) > 0 {
		c := newAddressConfig(registryAddress)
		url, err := c.toURL()
		if err != nil {
			panic(err)
		}
		RegistryCenter, err = extension.GetRegistry(c.getProtocol(), url)
		if err != nil {
			panic(err)
		}
	}
	if len(metadataReportAddress) > 0 {
		c := newAddressConfig(metadataReportAddress)
		url, err := c.toURL()
		if err != nil {
			panic(err)
		}
		fmt.Println(url.SubURL)
		factory := extension.GetMetadataReportFactory(c.getProtocol())
		MetadataReportCenter = factory.CreateMetadataReport(url)
	}
}

func newAddressConfig(address string) AddressConfig {
	config := AddressConfig{}
	config.Address = address
	var err error
	config.url, err = url.Parse(address)
	if err != nil {
		panic(err)
	}
	return config
}

// load database for mock rule storage, if dsn is empty, use sqlite in memory
func loadDatabaseConfig(dsn string) {
	var db *gorm.DB
	var err error
	if dsn == "" {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		panic(err)
	} else {
		DataBase = db
		// init table
		DataBase.AutoMigrate(&model.MockRuleEntity{})
	}
}