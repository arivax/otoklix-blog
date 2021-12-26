
# otoklix-blog

otoklix-blog adalah salah satu contoh aplikasi studi kasus dalam seleksi backend developer do otoklix.


## Installation & Configuration

Instalasi dan konfigurasi Golang dapat di lihat di https://go.dev/doc/install

Pada file otoklix-blog/src/.env set (terutama DB_NAME):

    APP_NAME=otoklix-blog
    HTTP_PORT=5300

    DB_NAME=/home/ariawan/Projects/clients/OTOKLIX/otoklix-blog/db/otoklix_blog

Pada file otoklix-blog/src/test/.env set (terutama DB_NAME):

    APP_NAME=otoklix-blog
    HTTP_PORT=5300

    DB_NAME=/home/ariawan/Projects/clients/OTOKLIX/otoklix-blog/src/test/otoklix_blog_test


## Usage

make run : digunakan untuk menjalankan aplikasi

make build : digunakan untuk build aplikasi di linux

make build-win : digunakan untuk build aplikasi di windows

make build-mac : digunakan untuk build aplikasi di macOS

make test : digunakan untuk test


## Examples

Get all:    curl -X GET  -L 'http://localhost:5300/otoklix-blog/api/BlogPost'

Get by ID:  curl -X GET -L 'http://localhost:5300/otoklix-blog/api/BlogPost/1'

Create:     curl -X POST -L 'http://localhost:5300/otoklix-blog/api/BlogPost' -H 'Content-Type: application/json' -d '{"title":"post-01","content":"content post-01"}'

Update:     curl -X PUT -L 'http://localhost:5300/otoklix-blog/api/BlogPost/4' -H 'Content-Type: application/json' -d '{"title":"post-04 ok","content":"content post-04 ok"}'

Delete:     curl -X DELETE -L 'http://localhost:5300/otoklix-blog/api/BlogPost/4' -H 'Content-Type: application/json' -d '{"title":"post-04 ok","content":"content post-04 ok"}'

