import { TestBed } from '@angular/core/testing';

import { CommentRoutesService } from './comment-routes.service';

describe('CommentRoutesService', () => {
  let service: CommentRoutesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CommentRoutesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
