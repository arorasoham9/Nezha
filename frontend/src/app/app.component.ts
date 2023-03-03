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

  constructor(private authService: SocialAuthService, private apiAuthService: AuthenticationService, private appService: AppsService ) { 
  }

  ngOnInit(): void {
    this.authService.authState.subscribe((user) => {
      this.user = user;
    }) 
  }

  signInWithGoogle(): void {
    this.authService.signIn(GoogleLoginProvider.PROVIDER_ID);

  }

  getToken(): void {
    console.log("Getting token.");
    console.log(this.user);
    OpenAPI.BASE = "http://localhost:8000"
    this.apiAuthService.login({email: this.user.email}).subscribe(body=>{
      OpenAPI.HEADERS = {"token": body};
      console.log(body)
      this.appService.getApps().subscribe(appList=>{
        console.log(appList)
      })
    })
  }


  getApps(): void {

  }

  signOut(): void {
    this.authService.signOut();
  }
}
