## task
- ソースと動作がイメージできる所から
- 階層によりprint結果はインデントをつける
- ディレクトリ指定がない場合は.。指定ありはそのディレクトリ検索
- `-d`ではディレクトリのみ検索する
- `-l`ではintチェックを行い、指定levelの階層のみ再帰的に検索する
- 正常な処理が完了してから出力を行う
    - 途中でエラー発生の可能性がある
- 存在しないディレクトリはエラー
- 色分けするpackage追加する

## tree command spec
```
#
# output sample
#

$ tree                                                                                                                                                                                                                         (git)-[master]
.
├── LICENSE
└── README.md

0 directories, 2 files
 
$tree --help
usage: tree [-acdfghilnpqrstuvxACDFJQNSUX] [-H baseHREF] [-T title ]
	[-L level [-R]] [-P pattern] [-I pattern] [-o filename] [--version]
	[--help] [--inodes] [--device] [--noreport] [--nolinks] [--dirsfirst]
	[--charset charset] [--filelimit[=]#] [--si] [--timefmt[=]<f>]
	[--sort[=]<name>] [--matchdirs] [--ignore-case] [--] [<directory list>]
  
  ------- Listing options -------
  -a            All files are listed.
  -d            List directories only.
  -f            Print the full path prefix for each file.
  -L level      Descend only level directories deep.

***

$ tree -L
tree: Missing argument to -L option.

$ tree -L 1
.
├── Applications
├── Boostnote
├── ChmodBPF
├── Desktop
├── Documents
├── Downloads
├── Dropbox
├── Fiddler2
├── Library
├── Movies
├── Music
├── PhpstormProjects
├── Pictures
├── Public
├── Sites
├── VirtualBox VMs
├── Workspace
├── download-mac
├── dwhelper
├── go
├── google-cloud-sdk
├── highlight.css
├── less
├── lrbr-web
├── mbox
├── rtf
└── vim_install.sh

21 directories, 6 files

***

% tree -L 1 -d
.
├── Applications
├── Boostnote
├── ChmodBPF
├── Desktop
├── Documents
├── Downloads
├── Dropbox
├── Fiddler2
├── Library
├── Movies
├── Music
├── PhpstormProjects
├── Pictures
├── Public
├── Sites
├── VirtualBox VMs
├── Workspace
├── dwhelper
├── go
├── google-cloud-sdk
└── lrbr-web

21 directories
```
