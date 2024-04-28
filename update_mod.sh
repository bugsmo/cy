#!bin/bash
          DIR="/home/murphy/code/cy/"

          for file in `find contrib -name go.mod`; do
              dirName=$(dirname $file)
            #   echo $DIR$dirName
              cd $DIR$dirName
              go get all
              go mod tidy
          done
          