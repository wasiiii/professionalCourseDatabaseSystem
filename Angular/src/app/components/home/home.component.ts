import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  constructor(
    public http : HttpClient,
    private router :Router
    ) { }

  ngOnInit() {
    this.initShow()
  }

  showManagement:boolean = false
  showClassTable:boolean = false
  showGradesRecord:boolean = false
  showCourseSelection:boolean = false

  initShow(){
    const httpOptions = {headers: new HttpHeaders({
      'Content-Type':'application/json'
    })};

    var role = window.localStorage.getItem('role')
    if(role == "0"){
      this.showClassTable = true
      this.showCourseSelection = true
    }
    else if(role == "1"){
      this.showGradesRecord = true
    }
    else if(role == "2"){
      this.showClassTable = true
      this.showGradesRecord = true
    }
    else if(role == "3"){
      this.showManagement = true
      this.showGradesRecord = true
    }
    else if(role == "4"){
      this.showClassTable = true
      this.showManagement = true
      this.showGradesRecord = true
    }
  }

}
