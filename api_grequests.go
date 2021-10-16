resp, err := grequests.post("http://api.jimmy.com/pants?type=skinnyjeans",
			&grequests.RequestOptions{
        JSON:   map[string]string{"token": "mah token", "boo":"baz", "bop":{"foo":{"bat":"boo"}}},
				Headers: map[string]string{"X-Riscosity-Tkn": token},
				IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
			})
