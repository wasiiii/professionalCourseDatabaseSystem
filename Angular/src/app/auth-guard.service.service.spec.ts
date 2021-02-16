import { TestBed } from '@angular/core/testing';

import { AuthGuard.ServiceService } from './auth-guard.service.service';

describe('AuthGuard.ServiceService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: AuthGuard.ServiceService = TestBed.get(AuthGuard.ServiceService);
    expect(service).toBeTruthy();
  });
});
