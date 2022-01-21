## go-libseal

## 简介

`go-libseal` 使用golang语言生成企业和个人印章。

## 功能

- 生成企业印章
- 生成个人印章

## 使用

1、获取包 `go get github.com/xwc1125/go-libseal`
2、生成企业印章调用

```
    dc, err := BuildCompanySeal("北京xxx科技有限公司", "110xxxxxxxx55", "Songti.ttc")
	if err != nil {
		t.Fatal(err)
	}
	os.MkdirAll("./img", os.ModePerm)
	err = dc.SavePNG("./img/out-company.png")
	if err != nil {
		t.Fatal(err)
	}
```

2、生成个人印章调用

```
	dc, err := BuildPersonalSeal("小明", "Songti.ttc")
    if err != nil {
        t.Fatal(err)
    }
    os.MkdirAll("./img", os.ModePerm)
    err = dc.SavePNG("./img/out-personal.png")
    if err != nil {
        t.Fatal(err)
    }
```

## 示例

- 企业印章

![企业印章](./img/out-company.png)

- 个人印章

![小明印章](./img/out-personal.png)
![小明印章](./img/out-personal2.png)

## 证书

`go-libseal` 的源码允许用户在遵循 [Apache 2.0 开源证书](LICENSE) 规则的前提下使用。

## 版权

Copyright@2022 xwc1125

![xwc1125](./logo.png)
