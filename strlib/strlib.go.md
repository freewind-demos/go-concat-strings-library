定义了一个名为`strlib`的package。其它地方使用该package下的函数时，必须使用以下方式导入：

```
import "github.com/golang-demos/go-concat-strings-library/strlib"
```

哪怕是在同一个项目中。

注意点
-----

### `package strlib`

当前package为`strlib`，它必须跟所在目录同名。

该目录下所有的go文件，都必须声明为同一个package，它们同属于一个package。

### `ConcatStrs`

以大写字母开头，表示它是一个public的函数，可以被别的包下的代码使用。如果是小写，则为private，只能在当前package中可见。
