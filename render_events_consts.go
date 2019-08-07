package main

func (renderer *Renderer) RenderEvents_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)
	f.WriteLine(`// Types used by the API.`)
	f.WriteLine(`const (`)
	for _,v := range renderer.Model.Types{
		if v.Name == "user"{
			f.WriteLine(`UsersFactStreamType = " users_facts "`)

		}
		if v.Name == "account"{
			f.WriteLine(`AccountsFactStreamType = " accounts_facts "`)

		}
	}
	f.WriteLine(`)`)
	f.WriteLine(`const (`)
	for _,v := range renderer.Model.Types{
		if v.Name == "user"{
			f.WriteLine(`UsersCommandsStreamType = " users_commands "`)

		}
		if v.Name == "account"{
			f.WriteLine(`AccountsCommandsStreamType = " accounts_commands "`)

		}
		if v.Name =="token" {
			f.WriteLine(`IdPCommandsStreamType = " idp_commands "`)
		}
	}
	f.WriteLine(`)`)
	f.WriteLine(`const (`)
	arrModelType := renderer.Model.Types
	for _,v := range arrModelType{
		//log.Print(v.TypeName)
		if v.TypeName == "CreateAccountUserResponses"{
			f.WriteLine(`UserCreatedFactType= " user_created "`)
		}
		if v.TypeName == "PatchUserResponses"{
			f.WriteLine(`UserPatchedFactType = " user_patched "`)

		}
		if v.TypeName == "DeleteUserResponses"{
			f.WriteLine(`UserDeletedFactType = " user_deleted "`)

		}
		if v.TypeName == "CreateAccountResponses"{
			f.WriteLine(`AccountCreatedFactType = " account_created "`)

		}
		if v.TypeName == "PatchAccountResponses"{
			f.WriteLine(`AccountPatchedFactType = " account_patched "`)

		}
		if v.TypeName == "DeleteUserResponses"{
			f.WriteLine(`AccountDeletedFactType = " account_deleted "`)

		}
	}
	f.WriteLine(`)`)
	f.WriteLine(`const (`)
	for _,v := range renderer.Model.Types{
		if v.TypeName == "GetUserParameters"{
			f.WriteLine(`CreateUserCommandType = " create_user "`)

		}
		if v.TypeName == "PatchUserParameters"{
			f.WriteLine(`PatchUserCommandType = " patch_user "`)

		}
		if v.TypeName == "DeleteUserParameters"{
			f.WriteLine(`DeleteUserCommandType = " delete_user "`)

		}
		if v.TypeName == "CreateAccountUserParameters"{
			f.WriteLine(`CreatedAccountCommandType = " create_account "`)

		}
		if v.TypeName == "PatchAccountParameters"{
			f.WriteLine(`UPatchAccountCommandType = " patch_account "`)

		}
		if v.TypeName == "PatchUserParameters"{
			f.WriteLine(`DeleteAccountCommandType = " delete_account "`)

		}
		if v.TypeName == "SignupResponses"{
			f.WriteLine(`SignupAccountCommandType = " signup_account "`)
		}
	}
	f.WriteLine(`)`)

	return f.Bytes(), nil
}