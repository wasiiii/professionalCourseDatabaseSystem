import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { fadeMotion } from 'ng-zorro-antd';

@Component({
  selector: 'app-show-course',
  templateUrl: './show-course.component.html',
  styleUrls: ['./show-course.component.scss']
})
export class ShowCourseComponent implements OnInit {

  constructor(public http:HttpClient, 
    private router:Router,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.getCourseInfo()
  }

  teaMap
  allCourse

  getCourseInfo(){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/getCourseInfo';
    this.http.post(api, {'account':account, 'role':role}, httpOptions).subscribe((response:any)=>{
      //console.log(response)
      if(response.check){
        this.teaMap = response.teaMap
        this.teaMap['无'] = "无"
        this.allCourse = response.allCourse
        for (var i = 0; i < this.allCourse.length; i++){
          this.allCourse[i].selectedTeaNameAdd = ""
          if(this.allCourse[i].Group.length > 0) this.allCourse[i].selectedTeaNameDel = this.allCourse[i].Group[0]
        }
        
      }
    });
  }
  teaChangeAdd(item, value){
    item.selectedTeaNameAdd = value
  }
  teaChangeDel(item, value){
    item.selectedTeaNameDel = value

  }

  add(item){
    if (item.selectedTeaNameAdd==""){
      this.createNotification("error", "警告", "请选择课程组外的教师")
      return
    }
    var flag = false
    for (var i = 0; i < item.Group.length; i++){
      if (item.selectedTeaNameAdd == item.Group[i]) {
        flag = true
        break
      }
    }
    if (flag){
      this.createNotification("error", "警告", "请选择课程组外的教师")
      return
    }
    
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/changeGroup';
    this.http.post(api, {'account':account, 'role':role,"acction":0, "courseID":item.Course_ID, "teaID":item.selectedTeaNameAdd}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        item.Group.push(item.selectedTeaNameAdd)
      }else{
        this.createNotification("error", "错误", "后台错误")
      }
    });
  }

  delete(item){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/changeGroup';
    this.http.post(api, {'account':account, 'role':role,"acction":1, "courseID":item.Course_ID, "teaID":item.selectedTeaNameDel}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        var key = item.Group.indexOf(item.selectedTeaNameDel)
        item.Group.splice(key, 1)
        if(item.Group.length > 0) item.selectedTeaNameDel = item.Group[0]
        else item.selectedTeaNameDel = ""
      }else{
        this.createNotification("error", "错误", "后台错误")
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
