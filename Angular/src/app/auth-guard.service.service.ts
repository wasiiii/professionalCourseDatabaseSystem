import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {

  constructor(
    private router : Router) { }
 
  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): boolean {
      var account = window.localStorage.getItem('account')
      var role = window.localStorage.getItem('role')
      if(!account || !role){
        this.router.navigateByUrl('login')
        return false
      }
      return true;
  }
}
