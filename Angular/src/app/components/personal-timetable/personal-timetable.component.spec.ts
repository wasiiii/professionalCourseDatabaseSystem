import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonalTimetableComponent } from './personal-timetable.component';

describe('PersonalTimetableComponent', () => {
  let component: PersonalTimetableComponent;
  let fixture: ComponentFixture<PersonalTimetableComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PersonalTimetableComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PersonalTimetableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
