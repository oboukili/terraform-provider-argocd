package argocd

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func applicationSpecSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		MinItems:    1,
		MaxItems:    1,
		Description: "ArgoCD App application resource specs. Required attributes: destination, source.",
		Required:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"destination": {
					Type:         schema.TypeSet,
					Required:     true,
					RequiredWith: []string{"source"},
					MinItems:     1,
					MaxItems:     1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"server": {
								Type:     schema.TypeString,
								Required: true,
							},
							"namespace": {
								Type:     schema.TypeString,
								Required: true,
							},
						},
					},
				},
				"source": {
					Type:         schema.TypeList,
					Required:     true,
					RequiredWith: []string{"destination"},
					MinItems:     1,
					MaxItems:     1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"repo_url": {
								Type:     schema.TypeString,
								Required: true,
							},
							"path": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"target_revision": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"chart": {
								Type:         schema.TypeString,
								RequiredWith: []string{"helm"},
							},
							"helm": {
								Type:         schema.TypeList,
								RequiredWith: []string{"chart"},
								MaxItems:     1,
								MinItems:     1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"value_files": {
											Type:     schema.TypeList,
											Optional: true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"values": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"parameters": {
											Type:     schema.TypeList,
											Optional: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:         schema.TypeString,
														RequiredWith: []string{"value"},
													},
													"value": {
														Type:         schema.TypeString,
														RequiredWith: []string{"name"},
													},
													"force_string": {
														Type:        schema.TypeBool,
														Description: "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
														Optional:    true,
													},
												},
											},
										},
										"release_name": {
											Type:        schema.TypeString,
											Description: "The Helm release name. If omitted it will use the application name",
											Optional:    true,
										},
									},
								},
							},
							"kustomize": {
								Type:     schema.TypeList,
								MaxItems: 1,
								MinItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name_prefix": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"name_suffix": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"version": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"images": {
											Type:     schema.TypeSet,
											Optional: true,
											Elem: schema.Schema{
												Type: schema.TypeString,
											},
										},
										"common_labels": {
											Type:         schema.TypeMap,
											Optional:     true,
											Elem:         &schema.Schema{Type: schema.TypeString},
											ValidateFunc: validateMetadataLabels,
										},
									},
								},
							},
							"ksonnet": {
								Type:     schema.TypeList,
								MaxItems: 1,
								MinItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"environment": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"parameters": {
											Type:     schema.TypeSet,
											Optional: true,
											Elem: schema.Resource{
												Schema: map[string]*schema.Schema{
													"component": {
														Type: schema.TypeString,
														RequiredWith: []string{
															"name",
															"value",
														},
													},
													"name": {
														Type: schema.TypeString,
														RequiredWith: []string{
															"value",
															"component",
														},
													},
													"value": {
														Type: schema.TypeString,
														RequiredWith: []string{
															"name",
															"component",
														},
													},
												},
											},
										},
									},
								},
							},
							"directory": {
								Type:     schema.TypeList,
								MaxItems: 1,
								MinItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"recurse": {
											Type:     schema.TypeBool,
											Optional: true,
										},
										"jsonnet": {
											Type:     schema.TypeList,
											Optional: true,
											Elem: schema.Resource{
												Schema: map[string]*schema.Schema{
													"ext_vars": {
														Type:     schema.TypeList,
														Optional: true,
														Elem: schema.Resource{
															Schema: map[string]*schema.Schema{
																"name": {
																	Type: schema.TypeString,
																	RequiredWith: []string{
																		"value",
																		"code",
																	},
																},
																"value": {
																	Type: schema.TypeString,
																	RequiredWith: []string{
																		"name",
																		"code",
																	},
																},
																"code": {
																	Type: schema.TypeString,
																	RequiredWith: []string{
																		"name",
																		"value",
																	},
																},
															},
														},
													},
													"tlas": {
														Type:     schema.TypeList,
														Optional: true,
														Elem: schema.Resource{
															Schema: map[string]*schema.Schema{
																"name": {
																	Type: schema.TypeString,
																	RequiredWith: []string{
																		"value",
																		"code",
																	},
																},
																"value": {
																	Type: schema.TypeString,
																	RequiredWith: []string{
																		"name",
																		"code",
																	},
																},
																"code": {
																	Type: schema.TypeString,
																	RequiredWith: []string{
																		"name",
																		"value",
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
							"plugin": {
								Type:     schema.TypeList,
								MaxItems: 1,
								MinItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"env": {
											Type:     schema.TypeSet,
											Optional: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:     schema.TypeString,
														Optional: true,
													},
													"value": {
														Type:     schema.TypeString,
														Optional: true,
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
				"project": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"sync_policy": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					MinItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"automated": {
								Type:     schema.TypeList,
								Optional: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"prune": {
											Type:     schema.TypeBool,
											Optional: true,
										},
										"self_heal": {
											Type:     schema.TypeBool,
											Optional: true,
										},
									},
								},
							},
							"sync_options": {
								Type:     schema.TypeSet,
								Optional: true,
								Elem: schema.Schema{
									Type: schema.TypeString,
									// TODO: add a validator
								},
							},
						},
					},
				},
				"ignore_differences": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"group": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"kind": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"name": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"namespace": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"json_pointers": {
								Type:     schema.TypeSet,
								Optional: true,
								Elem: schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
				"info": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"value": {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
				"revision_history_limit": {
					Type:     schema.TypeInt,
					Optional: true,
				},
			},
		},
	}
}
