import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-course-arranging',
  templateUrl: './course-arranging.component.html',
  styleUrls: ['./course-arranging.component.scss']
})
export class CourseArrangingComponent implements OnInit {

  constructor(public http:HttpClient,
    private notification: NzNotificationService) { }

  ngOnInit() {
  }

}
