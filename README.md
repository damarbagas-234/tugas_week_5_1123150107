# Test Sukses

<p align="center">
  <img width="857" height="482" alt="Screenshot (621)" src="https://github.com/user-attachments/assets/1e0e0526-8f5c-45ac-89e5-bd160eca19ca" />
<img width="968" height="563" alt="Screenshot (622)" src="https://github.com/user-attachments/assets/c615176b-a7f5-45b6-a861-c183237aeff7" />
<img width="951" height="535" alt="Screenshot (623)" src="https://github.com/user-attachments/assets/82ad7b89-ab8e-4cbd-b1e2-6ebdf5cfd455" />
<img width="961" height="586" alt="Screenshot (625)" src="https://github.com/user-attachments/assets/fc849225-f4a9-4acf-92d7-ac75e5468ecc" />

</p>

---
# Database


<p align="center">
<img width="857" height="482" alt="Screenshot (626)" src="https://github.com/user-attachments/assets/f1dc349d-61cb-491c-8f57-6d0f0f73a05a" />
<img width="951" height="535" alt="Screenshot (627)" src="https://github.com/user-attachments/assets/74ea661b-cbba-4227-b89e-7532df482c51" />
<img width="734" height="412" alt="Screenshot (628)" src="https://github.com/user-attachments/assets/f2b555aa-e8cf-4421-bfbc-0d5d6aa1101e" />
</p>

--- 

# Alur

<img width="6250" height="3165" alt="User Interaction Sequence-2026-03-16-044558" src="https://github.com/user-attachments/assets/30ec7c41-eca8-4bda-ba05-de9e49b9d67b" />

```bash
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB
Client->>Middleware: HTTP Request
Middleware->>Handler: forward request
Handler->>Service: CreateUser()
Service->>Repository: SaveUser()
Repository->>DB: INSERT USER
DB-->>Repository: OK
Repository-->>Service: user saved
Service-->>Handler: user response
Handler-->>Client: JSON Response
```

