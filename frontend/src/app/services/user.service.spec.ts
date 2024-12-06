import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { UserService, User } from './user.service';

describe('UserService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [UserService]
    });
    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  const mockUser: User = {
    user_id: 1,
    user_name: 'testuser',
    first_name: 'Test',
    last_name: 'User',
    email: 'test@example.com',
    user_status: 'A',
    department: 'IT'
  };

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get users', () => {
    const mockUsers = { data: [mockUser] };

    service.getUsers().subscribe(users => {
      expect(users).toEqual([mockUser]);
    });

    const req = httpMock.expectOne('http://localhost:8080/api/users');
    expect(req.request.method).toBe('GET');
    req.flush(mockUsers);
  });

  it('should get a single user', () => {
    service.getUser(1).subscribe(user => {
      expect(user).toEqual(mockUser);
    });

    const req = httpMock.expectOne('http://localhost:8080/api/user/1');
    expect(req.request.method).toBe('GET');
    req.flush(mockUser);
  });

  it('should add a user', () => {
    service.addUser(mockUser).subscribe(user => {
      expect(user).toEqual(mockUser);
    });

    const req = httpMock.expectOne('http://localhost:8080/api/user');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(mockUser);
    req.flush(mockUser);
  });

  it('should update a user', () => {
    service.updateUser(mockUser).subscribe(user => {
      expect(user).toEqual(mockUser);
    });

    const req = httpMock.expectOne(`http://localhost:8080/api/user/${mockUser.user_id}`);
    expect(req.request.method).toBe('PUT');
    expect(req.request.body).toEqual(mockUser);
    req.flush(mockUser);
  });

  it('should delete a user', () => {
    service.deleteUser(1).subscribe(response => {
      expect(response).toBeUndefined();
    });

    const req = httpMock.expectOne('http://localhost:8080/api/user/1');
    expect(req.request.method).toBe('DELETE');
  });

  it('should handle errors', () => {
    service.getUsers().subscribe({
      error: (error) => {
        expect(error.message).toBe('Username already exists');
      }
    });

    const req = httpMock.expectOne('http://localhost:8080/api/users');
    req.flush('Error', { 
      status: 409, 
      statusText: 'Conflict'
    });
  });
}); 