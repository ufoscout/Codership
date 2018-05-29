import { Component, OnInit } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthState, AuthStateModel } from '../auth/auth.state';
import { Login } from './login.model';
import { SetAuthData } from '../auth/auth.events';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styles: ['./header.component.css']
})
export class LoginComponent implements OnInit {

  @Select(AuthState) auth$: Observable<AuthStateModel>;
  model = new Login('', '');

  constructor(private store: Store, private http: HttpClient) {}

  ngOnInit(): void {
    this.auth$.subscribe(auth => {
      console.log("Auth data is: " + JSON.stringify(auth));
    });
  }

  submitLogin(): void {
    //this.store.dispatch(new SetAuthData({username: this.model.username, token: this.model.password}));
    this.http.post("/api/auth/login", {
      username: this.model.username,
      password: this.model.password,
    }).subscribe(response => {
      console.log("Response: " + JSON.stringify(response));
      this.store.dispatch(new SetAuthData({username: this.model.username, token: response['token']}));
    });
  }

  callPublic(): void {
    this.http.get("/api/auth/test/public").subscribe(response => {
      console.log("Response: " + JSON.stringify(response));
    });
  }

  callProtected(): void {
    this.http.get("/api/auth/test/protected").subscribe(response => {
      console.log("Response: " + JSON.stringify(response));
    });
  }

}
