1. generate model from parameters object, e. g.
	```yaml
	/signup:
        post:
          description: user signup process
          operationId: signup
          requestBody:
            content:
              application/x-www-form-urlencoded:
                schema:
                  type: object
                  required:
                    - company
                    - name
                    - email
                    - password
                    - newsSubscription
                  properties:
                    company:
                      maxLength: 255
                      type: string
                      description: Company the account belongs to
                    name:
                      maxLength: 255
                      type: string
                      description: first name and last name of user
                    email:
                      type: string
                      description: the email of the user
                      format: email
                    password:
                      type: string
                      description: password for the user
                      format: password
                    newsSubscription:
                      type: boolean
                      description: |
                        We periodically send feature updates, optimization recommendations, and in-depth tutorials.
                        Would you like to receive these?
            required: true
	```  

shoud generate struct as


  ```go 
type SignupParameters struct {
                          
                          	// Company the account belongs to
                          	Company string `json:"company"`
                          
                          	// first name and last name of user
                          	Name string `json:"name"`
                          
                          	// the email of the user
                          	Email string `json:"email"`
                          
                          	// password for the user
                          	Password string `json:"password"`
                          
                          	// We periodically send feature updates, optimization recommendations, and in-depth tutorials. Would you like to receive these? 
                          	NewsSubscription bool `json:"newsSubscription"`
                          }
  ```
2. client.go tries to return single entity instead of slice
3. if response contains just string generated client is not compiled

	```yaml
      responses:
        '200':
          description: password has been changed successful
          content:
            application/json:
              schema:
                type: string
	```
4. 