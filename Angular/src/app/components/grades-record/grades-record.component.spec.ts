import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GradesRecordComponent } from './grades-record.component';

describe('GradesRecordComponent', () => {
  let component: GradesRecordComponent;
  let fixture: ComponentFixture<GradesRecordComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GradesRecordComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GradesRecordComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
