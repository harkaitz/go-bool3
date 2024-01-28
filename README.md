# GO-BOOL3

This is a dirty library that defines a 3 state boolean. This is used
in HTML checkboxes.

## Go documentation

    package bool3 // import "github.com/harkaitz/go-bool3"
    
    type Bool string

## Go struct Bool

    package bool3 // import "."
    
    type Bool string
    
    func (b Bool) IsSet() bool
    func (b *Bool) Set(value bool)
    func (b Bool) String() string
    func (b Bool) Value() bool

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
