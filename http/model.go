package http

type CategoryTreeResponse struct {
	Categories []*Category `json:"categories"`
}

type Category struct {
	Id            int32       `json:"id"`
	Name          string      `json:"name"`
	Url           string      `json:"url"`
	SubCategories []*Category `json:"subCategories"`
}

var ExampleResponse = CategoryTreeResponse{
	Categories: []*Category{
		{
			Id:   1,
			Name: "Test",
			Url:  "/test/deneme",
			SubCategories: []*Category{
				{
					Id:   1,
					Name: "Test",
					Url:  "/test/deneme",
					SubCategories: []*Category{
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*Category{},
						},
						{
							Id:   1,
							Name: "Test",
							Url:  "/test/deneme",
							SubCategories: []*Category{
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*Category{},
								},
								{
									Id:   1,
									Name: "Test",
									Url:  "/test/deneme",
									SubCategories: []*Category{
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:   1,
											Name: "Test",
											Url:  "/test/deneme",
											SubCategories: []*Category{
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
											},
										},
									},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*Category{},
								},
							},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*Category{},
						},
					},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*Category{},
				},
			},
		},
		{
			Id:            1,
			Name:          "Test",
			Url:           "/test/deneme",
			SubCategories: []*Category{},
		},
		{
			Id:   1,
			Name: "Test",
			Url:  "/test/deneme",
			SubCategories: []*Category{
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*Category{},
				},
				{
					Id:   1,
					Name: "Test",
					Url:  "/test/deneme",
					SubCategories: []*Category{
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*Category{},
						},
						{
							Id:   1,
							Name: "Test",
							Url:  "/test/deneme",
							SubCategories: []*Category{
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*Category{},
								},
								{
									Id:   1,
									Name: "Test",
									Url:  "/test/deneme",
									SubCategories: []*Category{
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*Category{},
										},
										{
											Id:   1,
											Name: "Test",
											Url:  "/test/deneme",
											SubCategories: []*Category{
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*Category{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
