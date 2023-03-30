import {  GoogleLoginProvider, SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Component, OnInit } from '@angular/core';
import { AppsService, AuthenticationService, OpenAPI } from 'generated';
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
  appList: Array<String>;
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

  getToken(): void {
    console.log("Getting token.");
    console.log(this.user);
    OpenAPI.BASE = "http://localhost:8000"
    this.apiAuthService.login({Token: this.user.idToken}).subscribe(body=>{
      OpenAPI.HEADERS = {"token": body};
      console.log(body)
      this.appService.getApps().subscribe(appList=>{
        this.appList = appList;
        console.log(this.appList)
      })
    })
  }


  getApps(): void {

  }

  signOut(): void {
    this.authService.signOut();
  }
}
