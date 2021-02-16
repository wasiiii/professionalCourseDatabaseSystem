import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(public http:HttpClient, 
    private router:Router,
    private notification: NzNotificationService) { }

  ngOnInit() {
  }

  account: string = "";
  password: string = "";

  StuLogin(){
    if(this.account == "" || this.password == ""){
      this.createNotification('error', '警告', '用户名或密码不能为空')
      return
    }
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/stuLogin';
    this.http.post(api, {'stuID':this.account, 'password':this.password}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        this.createNotification('success', '成功', '登录成功')
        //本地记录账号
        window.localStorage.setItem('account', this.account)
        //本地记录角色，学生0，老师1，班主任2，系主任3，系主任+班主任4
        window.localStorage.setItem('role', "0")
        this.router.navigateByUrl("home")
      }
      else{
        this.createNotification('error', '警告', '用户名不存在或密码错误')
      }
    });
  }

  TeaLogin(){
    if(this.account == "" || this.password == ""){
      this.createNotification('error', '警告', '用户名或密码不能为空')
      return
    }
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/teaLogin';
    this.http.post(api, {'teaID':this.account, 'password':this.password}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        this.createNotification('success', '成功', '登录成功')
        //本地记录账号
        window.localStorage.setItem('account', this.account)
        //本地记录角色，学生0，老师1，班主任2，系主任3，系主任+班主任4
        var role = "1"
        if(response.type == 1 && response.isHeadmaster == 1) role = "4"
        if(response.type == 1 && response.isHeadmaster == 0) role = "3"
        if(response.type == 0 && response.isHeadmaster == 1) role = "2"
        window.localStorage.setItem('role', role)
        this.router.navigateByUrl("home")
      }
      else{
        this.createNotification('error', '警告', '用户名不存在或密码错误')
      }
    });
  }

  createNotification(type:string, title:string, text:string): void {
    this.notification.create(
      type,
      title,
      text
    );
  }
}
