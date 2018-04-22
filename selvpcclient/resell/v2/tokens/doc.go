/*
Package tokens provides the ability to create tokens through the Resell v2 API.

Example of creating a project-scoped token

  createOpts := tokens.TokenOpts{
    ProjectID: "f628616b452f4052b191161c26abba91",
  }
  token, err := tokens.Create(ctx, resellClient, createOpts)
  if err != nil {
    log.Fatal(err)
  }
	fmt.Println(token.ID)

Example of creating a domain-scoped token

  createOpts := tokens.TokenOpts{
    DomainName: "1122334455",
  }
  token, err := tokens.Create(ctx, resellClient, createOpts)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(token.ID)
*/
package tokens
