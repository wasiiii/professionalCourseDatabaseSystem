import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-grades-record',
  templateUrl: './grades-record.component.html',
  styleUrls: ['./grades-record.component.scss']
})
export class GradesRecordComponent implements OnInit {

  constructor(public http:HttpClient,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.gradesRecordInit()
  }

  allCourse
  courseMap
  selectedCourse
  showCourse

  gradesRecordInit(){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/gradesRecordInit';
    this.http.post(api, {'account':account, 'role':role}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        console.log(response)
        this.allCourse = response.allCourse
        this.courseMap = response.courseMap
        this.courseMap["全部"] = "全部"
        this.selectedCourse = "全部"
        this.showCourse = this.allCourse
        for (var i = 0; i < this.allCourse.length; i++){
          this.allCourse[i].newGrade = this.allCourse[i].Grade.Int32
        }
      }
      else{
        this.createNotification('error', '警告', '后台错误：'+response.err)
      }
    });
  }

  courseChange(value){
    this.showCourse = []
    if(value == "全部"){
      this.selectedCourse = "全部"
      this.showCourse = [] = this.allCourse
    }
    this.selectedCourse = this.courseMap[value]
    for (var i = 0; i < this.allCourse.length; i++){
      if (this.allCourse[i].CourseID == value) {
        this.showCourse.push(this.allCourse[i])
      }
    }
  }

  createNotification(type:string, title:string, text:string): void {
    this.notification.create(
      type,
      title,
      text
    );
  }

  commit(){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/commitGrades';
    this.http.post(api, {'account':account, 'role':role, 'allCourse':this.showCourse}, httpOptions).subscribe((response:any)=>{
        console.log(response)
      if(response.check){
        this.createNotification('success', '成功', '提交成功')
      }
      else{
        this.createNotification('error', '警告', '后台错误：'+response.err)
      }
    });
  }

}
