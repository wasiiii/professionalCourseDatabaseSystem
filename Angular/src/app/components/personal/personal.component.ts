import { GradesRecordComponent } from './../grades-record/grades-record.component';
import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-personal',
  templateUrl: './personal.component.html',
  styleUrls: ['./personal.component.scss']
})
export class PersonalComponent implements OnInit {

  constructor(public http:HttpClient, 
    private router:Router,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.getInformation()
  }

  roleName:string = ""
  name:string = ""
  class:string =""
  showClass:boolean = false
  headmaster:string = ""
  showHeadmaster:boolean = false
  gpa:number
  showGpa:boolean = false
  id:string = ""
  sex:string = ""
  showSex:boolean = false
  value:string = "是否为"
  rank

  getInformation(){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/getInformation';
    this.http.post(api, {'account':account, 'role':role}, httpOptions).subscribe((response:any)=>{
      if(response){
        if(response.check == false) {
          this.createNotification('error', '警告', '账号和角色不匹配')
        }
        if(role == "0"){
          this.roleName = "学生"
          this.name = response.stuName
          this.class = response.stuClass
          this.headmaster = response.classHeadmaster
          this.gpa = response.gpa
          this.id = response.stuID
          this.sex = response.stuSex
          this.rank = response.rank
          this.showClass = true
          this.showGpa = true
          this.showHeadmaster = true
          this.showSex = true
        }
        else if(role == "1"){
          this.roleName = "教师"
          this.name = response.teaName
          this.id = response.teaID
        }
        else if(role == "2"){
          this.roleName = "教师（班主任）"
          this.name = response.teaName
          this.class = response.class
          this.id = response.teaID
          this.showClass = true 
        }
        else if(role == "3"){
          this.roleName = "系主任"
          this.name = response.teaName
          this.id = response.teaID
        }
        else if(role == "4"){
          this.roleName = "系主任（班主任）"
          this.name = response.teaName
          this.class = response.class
          this.id = response.teaID
          this.showClass = true 
        }
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

  logout(){
    window.localStorage.removeItem('account');
    window.localStorage.removeItem('role');
    this.createNotification("error", "注销", "你已完成注销")
    this.router.navigateByUrl("login")
  }
}
