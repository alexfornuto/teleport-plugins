// Code generated by protoc-gen-tfschema. DO NOT EDIT.
// versions:
// 	protoc-gen-tfschema 0.0.1
// 	protoc        		v3.14.0
// source: types.proto

package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform/helper/schema"
)

// proto type fullname: types.RoleV3
func SchemaRoleV3() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"Kind": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"SubKind": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"Version": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"Metadata": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				// nested type name: types.Metadata
				Schema: map[string]*schema.Schema{
					"Name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"Namespace": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"Description": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"Labels": {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"Expires": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.IsRFC3339Time,
					},
					"ID": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"Spec": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				// nested type name: types.RoleSpecV3
				Schema: map[string]*schema.Schema{
					"Options": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							// nested type name: types.RoleOptions
							Schema: map[string]*schema.Schema{
								"ForwardAgent": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"MaxSessionTTL": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"PortForwarding": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: types.BoolValue
										Schema: map[string]*schema.Schema{
											"Value": {
												Type:     schema.TypeBool,
												Optional: true,
											},
										},
									},
								},
								"CertificateFormat": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"ClientIdleTimeout": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"DisconnectExpiredCert": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"BPF": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"PermitX11Forwarding": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"MaxConnections": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"MaxSessions": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"RequestAccess": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"RequestPrompt": {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
					"Allow": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							// nested type name: types.RoleConditions
							Schema: map[string]*schema.Schema{
								"Logins": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"Namespaces": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"NodeLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"Rules": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										// nested type name: types.Rule
										Schema: map[string]*schema.Schema{
											"Resources": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											"Verbs": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											"Where": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"Actions": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
										},
									},
								},
								"KubeGroups": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"Request": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: types.AccessRequestConditions
										Schema: map[string]*schema.Schema{
											"Roles": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											"ClaimsToRoles": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													// nested type name: types.AccessRequestClaimMapping
													Schema: map[string]*schema.Schema{
														"Claim": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"Value": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"Roles": {
															Type:     schema.TypeList,
															Optional: true,
															Elem:     &schema.Schema{Type: schema.TypeString},
														},
													},
												},
											},
											"Annotations": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues
													Schema: map[string]*schema.Schema{
														"Values": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.LabelValues.ValuesEntry
																Schema: map[string]*schema.Schema{
																	"key": {
																		Type:     schema.TypeString,
																		Optional: true,
																	},
																	"value": {
																		Type:     schema.TypeList,
																		Optional: true,
																		MaxItems: 1,
																		Elem: &schema.Resource{
																			// nested type name: wrappers.StringValues
																			Schema: map[string]*schema.Schema{
																				"Values": {
																					Type:     schema.TypeList,
																					Optional: true,
																					Elem:     &schema.Schema{Type: schema.TypeString},
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
								},
								"KubeUsers": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"AppLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"ClusterLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"KubernetesLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"DatabaseLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"DatabaseNames": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"DatabaseUsers": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
					"Deny": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							// nested type name: types.RoleConditions
							Schema: map[string]*schema.Schema{
								"Logins": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"Namespaces": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"NodeLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"Rules": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										// nested type name: types.Rule
										Schema: map[string]*schema.Schema{
											"Resources": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											"Verbs": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											"Where": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"Actions": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
										},
									},
								},
								"KubeGroups": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"Request": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: types.AccessRequestConditions
										Schema: map[string]*schema.Schema{
											"Roles": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											"ClaimsToRoles": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													// nested type name: types.AccessRequestClaimMapping
													Schema: map[string]*schema.Schema{
														"Claim": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"Value": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"Roles": {
															Type:     schema.TypeList,
															Optional: true,
															Elem:     &schema.Schema{Type: schema.TypeString},
														},
													},
												},
											},
											"Annotations": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues
													Schema: map[string]*schema.Schema{
														"Values": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.LabelValues.ValuesEntry
																Schema: map[string]*schema.Schema{
																	"key": {
																		Type:     schema.TypeString,
																		Optional: true,
																	},
																	"value": {
																		Type:     schema.TypeList,
																		Optional: true,
																		MaxItems: 1,
																		Elem: &schema.Resource{
																			// nested type name: wrappers.StringValues
																			Schema: map[string]*schema.Schema{
																				"Values": {
																					Type:     schema.TypeList,
																					Optional: true,
																					Elem:     &schema.Schema{Type: schema.TypeString},
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
								},
								"KubeUsers": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"AppLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"ClusterLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"KubernetesLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"DatabaseLabels": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues
										Schema: map[string]*schema.Schema{
											"Values": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.LabelValues.ValuesEntry
													Schema: map[string]*schema.Schema{
														"key": {
															Type:     schema.TypeString,
															Optional: true,
														},
														"value": {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																// nested type name: wrappers.StringValues
																Schema: map[string]*schema.Schema{
																	"Values": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem:     &schema.Schema{Type: schema.TypeString},
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
								"DatabaseNames": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"DatabaseUsers": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
	}
}

// proto type fullname: types.UserV2
func SchemaUserV2() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"Kind": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"SubKind": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"Version": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"Metadata": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				// nested type name: types.Metadata
				Schema: map[string]*schema.Schema{
					"Name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"Namespace": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"Description": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"Labels": {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"Expires": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.IsRFC3339Time,
					},
					"ID": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"Spec": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				// nested type name: types.UserSpecV2
				Schema: map[string]*schema.Schema{
					"OIDCIdentities": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							// nested type name: types.ExternalIdentity
							Schema: map[string]*schema.Schema{
								"ConnectorID": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"Username": {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
					"SAMLIdentities": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							// nested type name: types.ExternalIdentity
							Schema: map[string]*schema.Schema{
								"ConnectorID": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"Username": {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
					"GithubIdentities": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							// nested type name: types.ExternalIdentity
							Schema: map[string]*schema.Schema{
								"ConnectorID": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"Username": {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
					"Roles": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"Traits": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							// nested type name: wrappers.LabelValues
							Schema: map[string]*schema.Schema{
								"Values": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: wrappers.LabelValues.ValuesEntry
										Schema: map[string]*schema.Schema{
											"key": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"value": {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													// nested type name: wrappers.StringValues
													Schema: map[string]*schema.Schema{
														"Values": {
															Type:     schema.TypeList,
															Optional: true,
															Elem:     &schema.Schema{Type: schema.TypeString},
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
					"Status": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							// nested type name: types.LoginStatus
							Schema: map[string]*schema.Schema{
								"IsLocked": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"LockedMessage": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"LockedTime": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.IsRFC3339Time,
								},
								"LockExpires": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.IsRFC3339Time,
								},
							},
						},
					},
					"Expires": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.IsRFC3339Time,
					},
					"CreatedBy": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							// nested type name: types.CreatedBy
							Schema: map[string]*schema.Schema{
								"Connector": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: types.ConnectorRef
										Schema: map[string]*schema.Schema{
											"Type": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"ID": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"Identity": {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								"Time": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.IsRFC3339Time,
								},
								"User": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: types.UserRef
										Schema: map[string]*schema.Schema{
											"Name": {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
					"LocalAuth": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							// nested type name: types.LocalAuthSecrets
							Schema: map[string]*schema.Schema{
								"PasswordHash": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"TOTPKey": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"U2FRegistration": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										// nested type name: types.U2FRegistrationData
										Schema: map[string]*schema.Schema{
											"Raw": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"KeyHandle": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"PubKey": {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								"U2FCounter": {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
	}
}