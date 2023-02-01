package grpc

import "github.com/anilsenay/go-http-vs-grpc/grpc/api"

var ExampleResponse = api.CategoryTreeResponse{
	Categories: []*api.Category{
		{
			Id:   1,
			Name: "Test",
			Url:  "/test/deneme",
			SubCategories: []*api.Category{
				{
					Id:   1,
					Name: "Test",
					Url:  "/test/deneme",
					SubCategories: []*api.Category{
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*api.Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*api.Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*api.Category{},
						},
						{
							Id:   1,
							Name: "Test",
							Url:  "/test/deneme",
							SubCategories: []*api.Category{
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*api.Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*api.Category{},
								},
								{
									Id:   1,
									Name: "Test",
									Url:  "/test/deneme",
									SubCategories: []*api.Category{
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:   1,
											Name: "Test",
											Url:  "/test/deneme",
											SubCategories: []*api.Category{
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
											},
										},
									},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*api.Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*api.Category{},
								},
							},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*api.Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*api.Category{},
						},
					},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*api.Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*api.Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*api.Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*api.Category{},
				},
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*api.Category{},
				},
			},
		},
		{
			Id:            1,
			Name:          "Test",
			Url:           "/test/deneme",
			SubCategories: []*api.Category{},
		},
		{
			Id:   1,
			Name: "Test",
			Url:  "/test/deneme",
			SubCategories: []*api.Category{
				{
					Id:            1,
					Name:          "Test",
					Url:           "/test/deneme",
					SubCategories: []*api.Category{},
				},
				{
					Id:   1,
					Name: "Test",
					Url:  "/test/deneme",
					SubCategories: []*api.Category{
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*api.Category{},
						},
						{
							Id:            1,
							Name:          "Test",
							Url:           "/test/deneme",
							SubCategories: []*api.Category{},
						},
						{
							Id:   1,
							Name: "Test",
							Url:  "/test/deneme",
							SubCategories: []*api.Category{
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*api.Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*api.Category{},
								},
								{
									Id:            1,
									Name:          "Test",
									Url:           "/test/deneme",
									SubCategories: []*api.Category{},
								},
								{
									Id:   1,
									Name: "Test",
									Url:  "/test/deneme",
									SubCategories: []*api.Category{
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:            1,
											Name:          "Test",
											Url:           "/test/deneme",
											SubCategories: []*api.Category{},
										},
										{
											Id:   1,
											Name: "Test",
											Url:  "/test/deneme",
											SubCategories: []*api.Category{
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
												},
												{
													Id:            1,
													Name:          "Test",
													Url:           "/test/deneme",
													SubCategories: []*api.Category{},
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
