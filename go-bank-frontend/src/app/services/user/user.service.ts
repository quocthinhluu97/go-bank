import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { BehaviorSubject } from 'rxjs';
import { Router } from '@angular/router';


const httpOptions = {
    headers: new HttpHeaders({
        'Access-Control-Allow-Origin': '*',
    })
};

@Injectable({
    providedIn: 'root'
})

export class UserService {
    url: any = 'http://localhost:4200/api/';
    errorSubject: any = new BehaviorSubject<any>(null);
    errorMessage: any = this.errorSubject.asObservable();
    userSubject: any = new BehaviorSubject<any>(null);
    user: any = this.userSubject.asObservable();

    constructor(
        private http: HttpClient,
        private router: Router,
    ) { }

    login(Username: string, Password: string): any {
        this.http.post(`${this.url}login`, {Username, Password}, httpOptions).toPromise().then((res: any) => {
            if (res && res.jwt) {
                sessionStorage.setItem('jwt', res.jwt);
                this.errorSubject.next(null);

                if (res.data) {
                    this.userSubject.next(res.data);
                    sessionStorage.setItem('userId', res.data.ID)
                }
                this.router.navigateByUrl('');
            } else if (res.Message) {
                this.errorSubject.next(res.Message);
            }
        });
    }

    getUser() {
        const userId = sessionStorage.getItem('userId');
        const jwtToken = sessionStorage.getItem('jwt');
        const reqHeader = {
            headers: new HttpHeaders({
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + jwtToken,
            })
        };
        return this.http.get(`${this.url}user/${userId}`, reqHeader);
    }


    register(Username: string, Email: string, Password: string) {
        this.http.post(`${this.url}register`, { Username, Email, Password  }, httpOptions).toPromise().then((res: any) => {
            if (res && res.jwt) {
                sessionStorage.setItem('jwt', res.jwt);
                this.errorSubject.next(null);
                if (res.data) {
                    this.userSubject.next(res.data);
                    sessionStorage.setItem('userId', res.data.ID)
                }
                this.router.navigateByUrl('');
            } else if (res.Message) {
                this.errorSubject.next(res.Message);
            }
        });
    }
}
