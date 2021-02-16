import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowTeacherComponent } from './show-teacher.component';

describe('ShowTeacherComponent', () => {
  let component: ShowTeacherComponent;
  let fixture: ComponentFixture<ShowTeacherComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ShowTeacherComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ShowTeacherComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
