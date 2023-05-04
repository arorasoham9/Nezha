import {  GoogleLoginProvider, SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Component, OnInit, ɵɵresolveBody } from '@angular/core';
import { AppsService, AuthenticationService, OpenAPI, app, users_login_body } from 'generated';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [AuthenticationService, AppsService]
})
export class AppComponent implements OnInit {
  title = 'frontend';
  user: SocialUser;
  loggedIn: boolean;
  appList: Array<app>;
  userLogin: users_login_body;
  loading: boolean = false;
  private accessToken = '';

  constructor(private authService: SocialAuthService, private apiAuthService: AuthenticationService, private appService: AppsService ) { 
  }

  ngOnInit(): void {
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null)
      this.getToken()
    }) 
  }

  signInWithGoogle(): void {
    this.authService.signIn(GoogleLoginProvider.PROVIDER_ID);

  }

  connect(id: string) {
    this.appService.connectApp(id).subscribe(body => {
      console.log(body)
    })
  }

  getToken(): void {
    //console.log("Getting token.");
    //console.log(this.user);
    this.loading = true;
    OpenAPI.BASE = "http://localhost:8000"
    this.apiAuthService.login({Token: this.user.idToken}).subscribe(body=>{
      this.userLogin = JSON.parse(JSON.stringify(body));
      if (this.userLogin.Token) {
        OpenAPI.HEADERS = {"token": this.userLogin.Token};
        //console.log(body)
        this.appService.getApps().subscribe(appList=>{
          this.appList = appList;
          this.loading= false;
          //console.log(this.appList)
        })
      }
    })

  }


  getApps(): void {

  }

  signOut(): void {
    this.authService.signOut();
  }
}
