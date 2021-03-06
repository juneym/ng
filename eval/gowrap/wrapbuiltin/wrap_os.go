// Generated file, do not edit.

package wrapbuiltin

import (
	"reflect"

	"neugram.io/ng/eval/gowrap"

	wrap_os "os"
)

var pkg_wrap_os = &gowrap.Pkg{
	Exports: map[string]reflect.Value{

		"Args":              reflect.ValueOf(&wrap_os.Args).Elem(),
		"Chdir":             reflect.ValueOf(wrap_os.Chdir),
		"Chmod":             reflect.ValueOf(wrap_os.Chmod),
		"Chown":             reflect.ValueOf(wrap_os.Chown),
		"Chtimes":           reflect.ValueOf(wrap_os.Chtimes),
		"Clearenv":          reflect.ValueOf(wrap_os.Clearenv),
		"Create":            reflect.ValueOf(wrap_os.Create),
		"DevNull":           reflect.ValueOf(wrap_os.DevNull),
		"Environ":           reflect.ValueOf(wrap_os.Environ),
		"ErrClosed":         reflect.ValueOf(&wrap_os.ErrClosed).Elem(),
		"ErrExist":          reflect.ValueOf(&wrap_os.ErrExist).Elem(),
		"ErrInvalid":        reflect.ValueOf(&wrap_os.ErrInvalid).Elem(),
		"ErrNotExist":       reflect.ValueOf(&wrap_os.ErrNotExist).Elem(),
		"ErrPermission":     reflect.ValueOf(&wrap_os.ErrPermission).Elem(),
		"Executable":        reflect.ValueOf(wrap_os.Executable),
		"Exit":              reflect.ValueOf(wrap_os.Exit),
		"Expand":            reflect.ValueOf(wrap_os.Expand),
		"ExpandEnv":         reflect.ValueOf(wrap_os.ExpandEnv),
		"File":              reflect.ValueOf(reflect.TypeOf(wrap_os.File{})),
		"FileInfo":          reflect.ValueOf(reflect.TypeOf((*wrap_os.FileInfo)(nil)).Elem()),
		"FileMode":          reflect.ValueOf(reflect.TypeOf(wrap_os.FileMode(0))),
		"FindProcess":       reflect.ValueOf(wrap_os.FindProcess),
		"Getegid":           reflect.ValueOf(wrap_os.Getegid),
		"Getenv":            reflect.ValueOf(wrap_os.Getenv),
		"Geteuid":           reflect.ValueOf(wrap_os.Geteuid),
		"Getgid":            reflect.ValueOf(wrap_os.Getgid),
		"Getgroups":         reflect.ValueOf(wrap_os.Getgroups),
		"Getpagesize":       reflect.ValueOf(wrap_os.Getpagesize),
		"Getpid":            reflect.ValueOf(wrap_os.Getpid),
		"Getppid":           reflect.ValueOf(wrap_os.Getppid),
		"Getuid":            reflect.ValueOf(wrap_os.Getuid),
		"Getwd":             reflect.ValueOf(wrap_os.Getwd),
		"Hostname":          reflect.ValueOf(wrap_os.Hostname),
		"Interrupt":         reflect.ValueOf(&wrap_os.Interrupt).Elem(),
		"IsExist":           reflect.ValueOf(wrap_os.IsExist),
		"IsNotExist":        reflect.ValueOf(wrap_os.IsNotExist),
		"IsPathSeparator":   reflect.ValueOf(wrap_os.IsPathSeparator),
		"IsPermission":      reflect.ValueOf(wrap_os.IsPermission),
		"Kill":              reflect.ValueOf(&wrap_os.Kill).Elem(),
		"Lchown":            reflect.ValueOf(wrap_os.Lchown),
		"Link":              reflect.ValueOf(wrap_os.Link),
		"LinkError":         reflect.ValueOf(reflect.TypeOf(wrap_os.LinkError{})),
		"LookupEnv":         reflect.ValueOf(wrap_os.LookupEnv),
		"Lstat":             reflect.ValueOf(wrap_os.Lstat),
		"Mkdir":             reflect.ValueOf(wrap_os.Mkdir),
		"MkdirAll":          reflect.ValueOf(wrap_os.MkdirAll),
		"ModeAppend":        reflect.ValueOf(wrap_os.ModeAppend),
		"ModeCharDevice":    reflect.ValueOf(wrap_os.ModeCharDevice),
		"ModeDevice":        reflect.ValueOf(wrap_os.ModeDevice),
		"ModeDir":           reflect.ValueOf(wrap_os.ModeDir),
		"ModeExclusive":     reflect.ValueOf(wrap_os.ModeExclusive),
		"ModeNamedPipe":     reflect.ValueOf(wrap_os.ModeNamedPipe),
		"ModePerm":          reflect.ValueOf(wrap_os.ModePerm),
		"ModeSetgid":        reflect.ValueOf(wrap_os.ModeSetgid),
		"ModeSetuid":        reflect.ValueOf(wrap_os.ModeSetuid),
		"ModeSocket":        reflect.ValueOf(wrap_os.ModeSocket),
		"ModeSticky":        reflect.ValueOf(wrap_os.ModeSticky),
		"ModeSymlink":       reflect.ValueOf(wrap_os.ModeSymlink),
		"ModeTemporary":     reflect.ValueOf(wrap_os.ModeTemporary),
		"ModeType":          reflect.ValueOf(wrap_os.ModeType),
		"NewFile":           reflect.ValueOf(wrap_os.NewFile),
		"NewSyscallError":   reflect.ValueOf(wrap_os.NewSyscallError),
		"O_APPEND":          reflect.ValueOf(wrap_os.O_APPEND),
		"O_CREATE":          reflect.ValueOf(wrap_os.O_CREATE),
		"O_EXCL":            reflect.ValueOf(wrap_os.O_EXCL),
		"O_RDONLY":          reflect.ValueOf(wrap_os.O_RDONLY),
		"O_RDWR":            reflect.ValueOf(wrap_os.O_RDWR),
		"O_SYNC":            reflect.ValueOf(wrap_os.O_SYNC),
		"O_TRUNC":           reflect.ValueOf(wrap_os.O_TRUNC),
		"O_WRONLY":          reflect.ValueOf(wrap_os.O_WRONLY),
		"Open":              reflect.ValueOf(wrap_os.Open),
		"OpenFile":          reflect.ValueOf(wrap_os.OpenFile),
		"PathError":         reflect.ValueOf(reflect.TypeOf(wrap_os.PathError{})),
		"PathListSeparator": reflect.ValueOf(wrap_os.PathListSeparator),
		"PathSeparator":     reflect.ValueOf(wrap_os.PathSeparator),
		"Pipe":              reflect.ValueOf(wrap_os.Pipe),
		"ProcAttr":          reflect.ValueOf(reflect.TypeOf(wrap_os.ProcAttr{})),
		"Process":           reflect.ValueOf(reflect.TypeOf(wrap_os.Process{})),
		"ProcessState":      reflect.ValueOf(reflect.TypeOf(wrap_os.ProcessState{})),
		"Readlink":          reflect.ValueOf(wrap_os.Readlink),
		"Remove":            reflect.ValueOf(wrap_os.Remove),
		"RemoveAll":         reflect.ValueOf(wrap_os.RemoveAll),
		"Rename":            reflect.ValueOf(wrap_os.Rename),
		"SEEK_CUR":          reflect.ValueOf(wrap_os.SEEK_CUR),
		"SEEK_END":          reflect.ValueOf(wrap_os.SEEK_END),
		"SEEK_SET":          reflect.ValueOf(wrap_os.SEEK_SET),
		"SameFile":          reflect.ValueOf(wrap_os.SameFile),
		"Setenv":            reflect.ValueOf(wrap_os.Setenv),
		"Signal":            reflect.ValueOf(reflect.TypeOf((*wrap_os.Signal)(nil)).Elem()),
		"StartProcess":      reflect.ValueOf(wrap_os.StartProcess),
		"Stat":              reflect.ValueOf(wrap_os.Stat),
		"Stderr":            reflect.ValueOf(&wrap_os.Stderr).Elem(),
		"Stdin":             reflect.ValueOf(&wrap_os.Stdin).Elem(),
		"Stdout":            reflect.ValueOf(&wrap_os.Stdout).Elem(),
		"Symlink":           reflect.ValueOf(wrap_os.Symlink),
		"SyscallError":      reflect.ValueOf(reflect.TypeOf(wrap_os.SyscallError{})),
		"TempDir":           reflect.ValueOf(wrap_os.TempDir),
		"Truncate":          reflect.ValueOf(wrap_os.Truncate),
		"Unsetenv":          reflect.ValueOf(wrap_os.Unsetenv),
	},
}

func init() {
	if gowrap.Pkgs["os"] == nil {
		gowrap.Pkgs["os"] = pkg_wrap_os
	}
}
