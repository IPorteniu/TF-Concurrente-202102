import { TestBed } from '@angular/core/testing';

import { KnnService } from './knn;

describe('KnnService', () => {
  let service: KnnService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(KnnService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
