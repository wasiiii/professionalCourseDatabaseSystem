import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-show-student',
  templateUrl: './show-student.component.html',
  styleUrls: ['./show-student.component.scss']
})
export class ShowStudentComponent implements OnInit {

  constructor(public http:HttpClient, 
    private router:Router,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.getAllStu("")
    this.selectedClass = "全部"
    this.selectedTeaName = "无"
    this.makeShowStu(this.selectedClass)
  }
  class
  teacher
  allStu = []
  selectedClass:string
  teaName:string
  showStu = []
  selectedTeaName:string
  teacherID
  selectedTeaID:string

  getAllStu(value : string){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/getAllStu';
    this.http.post(api, {'account':account, 'role':role}, httpOptions).subscribe((response:any)=>{
      //console.log(response)
      if(response.check){
        this.class = response.class
        this.allStu = response.allStu
        this.showStu = response.allStu
        this.teacher = response.tea
        this.teacherID = response.teaIDs
        this.class["全部"] = "全部"
        this.teacher["无"] = "无"
      }
    });
  }

  classChange(value: string) {
    this.selectedClass = value;
    this.makeShowStu(this.selectedClass)
  }
  teaChange(value: string) {
    this.selectedTeaName = value;
    this.selectedTeaID = value
    //console.log(value)
  }

  makeShowStu(value:string){
    this.showStu = []
    if(value == "全部"){
      this.selectedTeaName = "无"
      for (var i = 0; i < this.allStu.length; i++){
        this.showStu.push(this.allStu[i])
      }
      return
    }else{
      this.selectedTeaID = this.class[value]
      this.selectedTeaName = this.selectedTeaID
      for (var i = 0; i < this.allStu.length; i++){
        if(this.allStu[i].StuClass == value){
          this.showStu.push(this.allStu[i])
        }
      }
    }
  }
  change(){
    //console.log(this.selectedClass)
    //console.log(this.selectedTeaID)
    if(this.selectedClass == "全部"){
      this.createNotification('error', '警告', '请选定班级')
      return
    }
    if(this.selectedTeaID == this.class[this.selectedClass]){
      this.createNotification('error', '警告', '请更改教师')
      return
    }

    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/changeClassHeadmaster';
    this.http.post(api, {'account':account, 'role':role, 'classID':this.selectedClass, 'teaID':this.selectedTeaID}, httpOptions).subscribe((response:any)=>{
      //console.log(response)
      if(response.check){
        this.createNotification("success", "成功", "更改成功")
      }else{
        this.createNotification("error", "警告", "更改失败")
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
