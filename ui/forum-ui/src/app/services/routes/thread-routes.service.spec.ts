import { TestBed } from '@angular/core/testing';

import { ThreadRoutesService } from './thread-routes.service';

describe('ThreadRoutesService', () => {
  let service: ThreadRoutesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ThreadRoutesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
