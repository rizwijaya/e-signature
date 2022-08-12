// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
//pragma abicoder v2;

contract Smartsign {
    address public sistem;
    address public userAccess;
    uint256 profileCount;
    uint256 signatureCount;
    uint256 documentCount;

    constructor() { 
        sistem = msg.sender;
        userAccess = msg.sender;
        profileCount = 1;
        signatureCount = 1;
        documentCount = 1;
    }

    //------Model data------//
    //Model User Profile
    struct Profile_public {
        uint256 profile_id;
        string idsignature;
        string password;
        string publickey;
        string role; // 1 admin, 2 user
    }

    struct Profile_private {
        address profile_address;
        string profile_id;
        string idsignature;
        string name;
        string identity_card;
        string email;
        string phone;
        string dateregistered;
    }
    //Model Signature Data
    struct Signature {
        uint256 signature_id;
        string profile_id;
        string signature;
        string datecreated;
    }
    struct Allsigners {
        address signer_address;
        uint256 signer_id;
    }
    //Model Document Data
    struct Document { //Document selama proses ttd dari awal sampai selesai
        uint256 document_id;
        address signers;
        uint256 signing; //Yang telah ttd
        string namefile;
        bytes32 hash_original; //identifier hash original file untuk get data history ttd dan urutkan paling baru
        string hash_file;
        string metadata;
        string file;
        string status; //1 No signed, 2 Request Signed, 3 waiting signed, 4 Signed
        uint256 totalsigned;
        string datesigned;
        uint256[] signerlist;
        mapping(uint256 => Allsigners ) allsigners;
        //address signers; //Alamat Signatures
    }
    //Model Document Signed Data
    // struct DocumentSigned { //Document terakhir saat berhasil di ttd
    //     uint256 document_signed_id;
    //     bytes32 hash_file;
    //     string metadata;
    //     string file;
    //     mapping (address => uint256) signers; //array yang menyimpan data penandatangannya.
    //     string datefinishing;
    // }
    //------End Model data------//

    //------Permission Access------//
    // Permission hanya admin/sistem yang dapat mengakses
    modifier onlySistem() {
        require(msg.sender == sistem);
        _;
    }
    //Permission User yang dapat mengakses
    modifier permissionUser() {
        require(msg.sender == userAccess);
        _;
    }
    //------End Permission Access------//

    //------Profile Function------//
    //------Create User Profile------//
    mapping(uint256 => Profile_public) public profiles_public;
    //First Register Before Private Key
    function add_profilefirst (string memory _idsignature, string memory _password, string memory _publickey, string memory _role) public onlySistem {
        Profile_public memory newProfile = 
            Profile_public({
                profile_id: profileCount,
                idsignature: _idsignature,
                password: _password,
                publickey: _publickey,
                role: _role
            });
        profiles_public[profileCount] = newProfile;
        profileCount += 1;
    }

    address[] public profile_private;
    mapping(address => Profile_private) public Profiles_private;
    //Second Step Register After Private Key and Save data in Key User
    function add_profile (address _address, string memory _profile_id, string memory _idsignature, string memory _name, string memory _identity_card, string memory _email, string memory _phone, string memory _dateregistered) public onlySistem {
            Profile_private memory newProfile_private = 
                Profile_private({
                    profile_address: _address,
                    profile_id : _profile_id,
                    idsignature: _idsignature,
                    name: _name,
                    identity_card: _identity_card,
                    email: _email,
                    phone: _phone,
                    dateregistered: _dateregistered
                });
            Profiles_private[_address] = newProfile_private;
            profile_private.push(_address);
    }
    //------End Create User Profile------//

    //Get Total Profile
    function get_totalprofile() public view returns (uint256) {
        return profileCount;
    }
    //------End Profile Function------//

    //------Signature Function------//
    //------Create Signature------//
    mapping(uint256 => Signature) public signatures;
    function add_signature(string memory _profile_id, string memory _signature, string memory _datecreated) public permissionUser {
        Signature memory newSignature =
            Signature({
                signature_id: signatureCount,
                profile_id: _profile_id,
                signature: _signature,
                datecreated: _datecreated
            });
        signatures[signatureCount] = newSignature;
        signatureCount += 1;
    }
    //------Get Signature------//
    function get_totalsignature() public view returns (uint256) {
        return signatureCount;
    }
    //------End Signature Function------//

    //------Document Function------//
    //------Signing Document------//
    mapping(bytes32 => Document) documents;
    bytes32[] documentlist;
    function createdstatus(bytes32 _hash) view internal returns (bool) {
        return (keccak256(bytes(documents[_hash].status)) != keccak256(bytes("1"))); //Document Not Exist
    }
    
    function signDoc(bytes32 _hash_original, string memory _namefile, string memory _hash_file, string  memory _metadata, string memory _file, string memory _status, uint256 _totalsigned, string memory _datesigned, uint256 signing) public {
        require(existdocument(_hash_original), "Document not found/Signing all");
        require(checkpermission(_hash_original, _totalsigned), "You are not allowed to sign this document");
        if (createdstatus(_hash_original)) { //Document Status Not Created
            documents[_hash_original].signing = signing-1; //Saat ini process ttd maka kurangi kondisi yang sudah dittd
        } else { //Document Status Created
            documents[_hash_original].signing = _totalsigned;
        }

        if(documents[_hash_original].signing-1 == 0) { //User terakhir yang melakukan ttd
            documents[_hash_original].status = _status;
        } else {
            documents[_hash_original].status = "4"; //Status Request Signed
        }

        documents[_hash_original].document_id = documentCount;
        documents[_hash_original].signers = msg.sender;
        documents[_hash_original].namefile = _namefile;
        documents[_hash_original].hash_original = _hash_original;
        documents[_hash_original].hash_file = _hash_file;
        documents[_hash_original].metadata = _metadata;
        documents[_hash_original].file = _file;
        documents[_hash_original].totalsigned = _totalsigned;
        documents[_hash_original].datesigned = _datesigned;
        documentlist.push(_hash_original);
        documentCount += 1;
    }

    //Add signers in hash_original
    function addsigners(bytes32 _hash_original, address _signers, uint256 _signer_id, uint256 total) public {
        require(existdocument(_hash_original), "Document not found/Signing All");
        require(checkpermission(_hash_original, total), "You are not allowed to sign this document");
        documents[_hash_original].signerlist[_signer_id] = _signer_id;
        documents[_hash_original].allsigners[_signer_id] = Allsigners({
            signer_address: _signers,
            signer_id: _signer_id
        });
    }
    function existdocument(bytes32 _hash) view internal returns (bool) {
        if (keccak256(bytes(documents[_hash].status)) != keccak256(bytes(""))) { //Document Not Exist
            return false;
        }
        if(documents[_hash].signing == 0) { //Telah melakukan ttd semua
            return false;
        }
        return true;
    }
    function checkpermission(bytes32 _hash, uint256 total) view internal returns (bool) {
        if (createdstatus(_hash)) { //Document Status Not Created
            for (uint256 i = 0; i < total; i++) {
                if (documents[_hash].allsigners[i].signer_address == msg.sender) {
                    return true;
                }
            }
        } else { //Document Status Created
            if(msg.sender == userAccess) {
                return true;
            }
            return false;
        }

        return false;
    }
    //------End Final Signing Document------//
    //------Get Document Latest------//

    //------End Get Document Latest------//
    //------Verify Document Function------//
    function verify_document(bytes32 _hash) public view returns (bool) {
        if(keccak256(bytes(documents[_hash].file)) != keccak256(bytes(""))) { //Document Exist
            return true;
        }
       return false;
    }
    // ------End Document Function------//

    //------Profile Function For Development Only (in case publickey and privatekey generated) ------//

    //------End Profile Function For Development Only (in case publickey and privatekey generated) ------//
}

    //------Get User Profile Login/Register Algoritma------//
    //Login User First
    //Get Profile Public By profile_id, idsignature, password, publickey // sistem permission
    //After Valid Login (with public key) -> set auth with privatekey and session (posisi sekarang didompet user)
    //Get Profile Private, All Data //User permission
    
    //Algoritma Signature rame"
    //User melakukan invite ke pihak yang akan menandatangani via email.
        //Pada blockchain akan memberikan ke masing" alamat profile dan membatasi aksesnya hanya diprofile tersebut
    //Kemudian masuk ke email masing" dan difitur permintaan (data diblockchain)
    //saat user membuka link tersebut, maka diminta login. Berikutnya jika tidak ada didompet maka doc tidak valid
    //tapi jika ada maka user dapat melakukan ttd,
    //untuk algo diblockchain akan melakukan generate semua data ttd pada document tersebut (by id/unique)
    //selanjutnya filter untuk document terbaru (fungsi filter adalah supaya pdf selalu update terbaru. agar user dapat melakukan ttd bergantian dengan document yang diupdate, supaya tidak perlu dimerge juga)
    //nantinya saat ada ttd, maka countnya didocument juga dilakukan berdasarkan jumlah penandatangannya
    //apabila countnya sudah terpenuhi (semua sudah ttd)
    //pdf terakhir adalah pdf jadi dan dikirimkan ke masing" tanda tangan
    //upload ke penyimpanan document signed beserta hash terakhirnya
    //fungsi verifikasi akan ke document signed dan memastikan hash beserta metadatanya merupakan metadata yang sama
    //kalau tidak sama maka tidak valid, begitu sebaliknya

