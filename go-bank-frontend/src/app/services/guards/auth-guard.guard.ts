import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { UserService } from '../user/user.service';

@Injectable({
    providedIn: 'root'
})
export class AuthGuardGuard implements CanActivate {
    constructor (
        private login: UserService,
    ) {}

    canActivate(): boolean {
        if (!this.login.isAuthenticated()) {
            return false;
        }
        return true
    }


}
