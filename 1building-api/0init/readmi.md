

##1 first we have to iniate the porour poject in respective dir
by using 
    go mod init github.com/<username>/<reponame>
    like 
    go mod init github.com/tejasthonge/<repo> 

    //this above repo is the what we want 
    //due to that it will create go.mod file 
        -this is look like bellow
            module github.com/tejasthonge/<repo>

            go 1.23.4

    //by using this,what are packages are used in our porject this will geted from this above module path

 
2## creating floder strucure

    ProjectFolder
        -cmd
            -<projectnam e> 
                -main.go
        -go.mod



3## creating .gitignore file 
    - we can create this manully but we use the third party extention
    - that is gitignore
        - pressing the cmd+shipt+p 
            -search for the  gitignore
            - select the language and inter
    - by this way it can automatically genarate the .gitignore file


