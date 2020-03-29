# SETUP

auto import for vscode https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code

* Encountering "You are neither in a module nor in your GOPATH" when importing local package
Use the go module.
Ref: https://github.com/golang/go/wiki/Modules#how-to-install-and-activate-module-support
     https://github.com/microsoft/vscode-go/issues/3086

* Episode 3
Package Json for marchalling the struct
Use of field tag

When a request comes into the server, go will call the handler's ServeHTTP function (must be implemented)

* Episode 4

```
d, err := json.Marshal(lp)
rw.Write(d)

e := json.NewEncoder(w)
return e.Encode(p)
```

