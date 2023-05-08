package handlers

// we need command args - we can monitor, or clone

// if clone we setup a new project, based on a repo tag or release - or latest
// we need to gather the data to do this.  project folder, tag

// if monitor - and not alreadt collected - we will gather data to log in to morpheus API
// if we have it, we will ask if we should used saved data, or create new connection
// we will monitor error logs
// we will detect changes in project folder, build the plugin and upload to morpheus using API
// we will confirm plugin appears in integrations/plugins

//https://github.com/go-survey/survey

//log.Fatalln("end")

/*	url := internal.PROJECT_URL

	// Replace "v1.0.0" with the tag you want to check out, or use "main" to check out the latest commit on the main branch.
	ref := internal.Tags["v0.0.1"] // or set defaut "main"
	folder := internal.DEFAULT_PROJECT_NAME

	// Create a directory to clone the repository into.
	err := os.Mkdir(folder, 0755)
	if err != nil {
		panic(err)
	}

	// Clone the repository into the directory.
	_, err = git.PlainClone(folder, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		//ReferenceName: plumbing.ReferenceName("refs/tags/" + ref),
		ReferenceName: plumbing.ReferenceName(""),
	})

	if err != nil {
		panic(err)
	}

	// remove .git folder
	/*gitFolder := filepath.Join(folder, ".git")
	if err := os.RemoveAll(gitFolder); err != nil {
		fmt.Println("could not remove git folder")
	}

	// Change into the repository directory.
	err = os.Chdir(folder)
	if err != nil {
		panic(err)
	}*/
