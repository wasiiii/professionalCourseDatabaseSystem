import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-course-selection',
  templateUrl: './course-selection.component.html',
  styleUrls: ['./course-selection.component.scss']
})
export class CourseSelectionComponent implements OnInit {

  constructor(public http:HttpClient,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.courseSelectInit()
  }

  course = []
  isTime:boolean

  courseSelectInit(){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/courseSelectInit';
    this.http.post(api, {'account':account, 'role':role}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        console.log(response)
        if(response.isTime){
          this.course = response.table
          this.isTime = true
        }else{
          this.isTime = false
          this.createNotification('error', '警告', '未到开课时间')
        }
      }
      else{
        this.createNotification('error', '警告', '后台错误：'+response.err)
      }
    });
  }

  choose(item){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/courseSelect';
    this.http.post(api, {'account':account, 'role':role, 'courseID':item.CourseID, 'isChosed': item.IsChosed}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        if(response.isTime){
          this.isTime = true
          if(item.IsChosed){
            item.StuNum--
          }else{
            item.StuNum++
          }
          item.IsChosed = ! item.IsChosed
        }else{
          this.isTime = false
          this.createNotification('error', '警告', '不是选课时间')
        }
      }
      else{
        this.createNotification('error', '警告', '后台错误：'+response.err)
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
