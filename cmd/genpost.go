package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"hidevops.io/hiboot/pkg/system"
	"hidevops.io/hiboot/pkg/app/cli"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hiboot/pkg/utils/io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var categoryTemplate = `---
title: ${title}
date: ${date}
weight: ${weight}
---
`

var postTemplate = `---
desc: 由 genpost (https://github.com/hidevopsio/genpost) 代码生成器生成
title: ${title}
date: ${date}
author: ${author}
---

## 子标题 1

正文...

## 子标题 2

正文...

`

// RootCommand is the root command
type GenPostCommand struct {
	cli.RootCommand

	category bool
}


// NewRootCommand the root command
func NewRootCommand() *GenPostCommand {
	c := new(GenPostCommand)
	c.Use = "genpost"
	c.Short = "genpost command"
	c.Long = "Run genpost command"
	c.Example = `
1. 生成类别

✔ 目录: articles
✔ 标题: 分享文章
✔ 排序 : 1

2. 生成文章
Use the arrow keys to navigate: ↓ ↑ → ←

? 选择类型:
    Hiboot云原生应用框架
  ▸ 分享文章
    代码阅读

✔ 标题 : 我的文章标题█
✔ 作者 : 邓冰寒
`

	flags := c.PersistentFlags()
	flags.BoolVarP(&c.category, "category", "c", false, "--category=true or -c")

	return c
}

// Run root command handler
func (c *GenPostCommand) Run(args []string) error {
	//log.SetLevel(log.DebugLevel)
	log.Infof("handle genpost command")

	root := filepath.Join(io.GetWorkDir(), "/content/")

	log.Debug(root)

	if c.category {
		c.genCategory(root)
	} else {
		c.genPost(root)
	}

	return nil
}

func (c *GenPostCommand) genPost(root string) error {
	var items []string
	var paths []string

	var files []string
	filepath.Walk(root, io.Visit(&files))
	for _, file := range files {
		if strings.Contains(file, "_index.md") {
			log.Debugf("file : %v", file)
			prop, err := system.ReadYamlFromFile(file)
			if err == nil {
				title, ok := prop["title"]
				if ok {
					items = append(items, title.(string))
					paths = append(paths, io.BaseDir(file))
				}
			}
		}
	}

	sel := promptui.Select{
		Label: "选择类型 ",
		Items: items,
	}

	idx, result, err := sel.Run()

	if err != nil {
		log.Debugf("Prompt failed %v\n", err)
	} else {
		log.Debugf("You choose %v: %q, %v\n", idx, result, paths[idx])
	}

	tt := promptui.Prompt{
		Label: "标题",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("不允许为空")
			}
			return nil
		},
	}
	title, _ := tt.Run()
	postTemplate = strings.Replace(postTemplate, "${title}", title, -1)

	input(&postTemplate,"作者 ", "${author}")

	t := time.Now()

	postTemplate = strings.Replace(postTemplate, "${date}", t.Format(time.RFC3339), -1)

	//fmt.Print(template)

	filename := fmt.Sprintf("%d-%02d-%02dT%02d%02d%02d.md",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	io.WriterFile(paths[idx], filename, []byte(postTemplate))

	fmt.Printf("已经生成文件：%v\n", filepath.Join(paths[idx], filename))

	return nil
}

func (c *GenPostCommand) genCategory(root string) error {

	f := promptui.Prompt{
		Label: "目录 ",
		Validate: func(input string) error {
			ok := isValid(input)
			if !ok {
				return errors.New("请输入有效的文件夹名称")
			}
			return nil
		},
	}
	folder, _ := f.Run()

	tt := promptui.Prompt{
		Label: "标题 ",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("不允许为空")
			}
			return nil
		},
	}
	title, _ := tt.Run()
	categoryTemplate = strings.Replace(categoryTemplate, "${title}", title, -1)

	w := promptui.Prompt{
		Label: "排序 ",
		Validate: func(input string) error {
			if _, err := strconv.ParseInt(input,10,64); err == nil {
				fmt.Printf("%q looks like a number.\n", input)
				return nil
			}
			//log.Error(input)
			return errors.New("无效输入，请输入数字")
		},
	}
	weight, _ := w.Run()
	postTemplate = strings.Replace(categoryTemplate, "${weight}", weight, -1)

	t := time.Now()
	postTemplate = strings.Replace(categoryTemplate, "${date}", t.Format(time.RFC3339), -1)

	path := filepath.Join(root, folder)

	filename := "_index.md"
	fullPath := filepath.Join(path, filename)
	log.Debugf(fullPath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		io.WriterFile(path, filename, []byte(categoryTemplate))
		fmt.Printf("已经生成文件：%v\n", fullPath)
	} else {
		fmt.Printf("文件 %v 已存在! \n", fullPath)
		return errors.New("文件已存在！")
	}

	return nil
}

func isValid(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}

	return false
}

func input(template *string, label string, varName string) (retVal string, err error) {
	title := promptui.Prompt{
		Label: label,
	}
	retVal, err = title.Run()

	*template = strings.Replace(*template, varName, retVal, -1)
	return
}