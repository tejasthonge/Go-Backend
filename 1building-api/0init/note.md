

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


4## seting the config file
    - it is mor emportant for production code 
        -- for that we will creating or writion code for that like show in 1student-api folder
    - setting the configartiow hase two way
        -- 1 setting envarmetal vairable 
        -- 2 file base conig 
    - for porductution code we use the file base config 
        -- due to that we can cotrole the versione essyly 
    
    - now create the config 
        - make config file in project root dir
            - inside that create local.ymal file

    now floder structurle loock like ->
        ProjectFolder
            -config
                -local.yaml
            -cmd
                -<projectnam e> 
                    -main.go
            -go.mod

    -- NOW THIS local.yam for the develpent 
        -for production we can create the production.yaml
    
5## now all the cofig data present in .yaml we have to serilise in go data struct to access them 
    --for that 
        -create interl or pkg folder in root of project and insid that create config folder and inside that create the config.go
        
        now porject structure will look like -->
            ProjectFolder
            -config
                -local.yaml
            -pkg
                -config
                    -config.go
            -cmd
                -<projectnam e> 
                    -main.go
            -go.mod


6## now for the searialize the data in config.go 
    -- ther diffreant way 
        - but we use the packgae for tha 
        that is clean env get form --> https://github.com/ilyakaznacheev/cleanenv

        add thes package using comad 
            go get -u github.com/ilyakaznacheev/cleanenv
    
    -- now this is dowinladed and new file is created name as go.sum 
    -- and also go.mod file also some changes

    now add the back tick on or cofig.go insde cofig struct as show in 1student-api 
    