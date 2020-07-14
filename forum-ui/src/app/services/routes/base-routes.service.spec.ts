import { TestBed } from '@angular/core/testing';

import { BaseRoutesService } from './base-routes.service';

describe('BaseRoutesService', () => {
  let service: BaseRoutesService<any>;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BaseRoutesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
