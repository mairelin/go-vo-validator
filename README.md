# go-vo-validator

This is a package for validating data received in a struct type. If you are working in basic validations on the server side may this is for you, v you only need to add some tags to the struct declaration then call the method to validate of this packages.
In this version the tags that are supported are:

- mandatory
- validateMin
- validateMax

This is an example of usage:

```
type VOExample struct {
	Name string `mandatory:"true"`
	Rating int  `validateMax:"5"`
}

go-vo-validator.Validate(&vo)
```




 
  