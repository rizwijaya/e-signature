// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
pragma abicoder v2;

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
    struct Profile {
        address profile_address;
        uint256 profile_id;
        string name;
        string email;
        string phone;
        string identity_card;
        string dateregistered;
        bool exist;
        //uint256[] signaturelist;
        mapping(uint256 => Signatures) signature;
    }
    //Model Signature Data
    struct Signatures {
        uint256 signature_id;
        string signature;
        string datecreated;
        bool exist;
    }

    //Model Document Data
    struct Document {
        //Document selama proses ttd dari awal sampai selesai
        uint256 document_id;
        bytes32 hash_original; //identifier hash original file untuk get data history ttd dan urutkan paling baru
        address signers;
        uint signing; //Yang telah ttd
        string namefile;
        string hash_file;
        string metadata;
        string hash_ipfs;
        uint state; //1 Process Signed, 2 Signed
        uint totalsigned;
        uint createdtime;
        uint completedtime;
        bool exist;
        mapping(address => Allsigners) allsigners;
    }
    //Model Document Signers
    struct Allsigners {
        address signer_address;
        bool state; //True is signed and False is not signed
        string date; //Tanggal TTD
        bool exist;
    }
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
    //Multipermissions
    mapping(address => bool) members;

    function addPermission(address _member) public {
        members[_member] = true;
    }

    function permission(address _member) internal view returns (bool) {
        if (members[_member]) {
            return true;
        }
        return false;
    }

    //Document Permission for Signing
    function docPermission(bytes32 _hash) internal view returns (bool) {
        if (documents[_hash].exist == false) {
            return false;
        }
        if (documents[_hash].signing == 0 || documents[_hash].state == 2) {
            //Tidak ada dokumen atau Telah melakukan ttd semua
            return false;
        }
        //Permission Address Signed or Not
        if (documents[_hash].allsigners[msg.sender].exist == true) {
            //Terdapat Address
            if (documents[_hash].allsigners[msg.sender].state == false) {
                //Jika address belum ttd
                return true;
            }
        }
        return false;
    }

    //------End Permission Access------//

    //------Profile Function------//
    //------Create User Profile------//
    address[] public profile;
    mapping(address => Profile) public Profiles;

    //Second Step Register After Private Key and Save data in Key User
    function add_profile(
        string memory _name,
        string memory _identity_card,
        string memory _email,
        string memory _phone,
        string memory _dateregistered
    ) public permissionUser {
        Profiles[msg.sender].profile_address = msg.sender;
        Profiles[msg.sender].profile_id = profileCount;
        Profiles[msg.sender].name = _name;
        Profiles[msg.sender].email = _email;
        Profiles[msg.sender].phone = _phone;
        Profiles[msg.sender].identity_card = _identity_card;
        Profiles[msg.sender].dateregistered = _dateregistered;
        Profiles[msg.sender].exist = true;
        profile.push(msg.sender);
        profileCount++;
    }

    //------Signing Profiles------//
    function profilesign(
        uint256 _sign_id,
        string memory sign,
        string memory timestamp
    ) public {
        Profiles[msg.sender].signature[_sign_id] = Signatures({
            signature_id: _sign_id,
            signature: sign,
            exist: true,
            datecreated: timestamp
        });
    }

    //------End Create User Profile------//
    //GetAllSign

    //Get Total Profile
    function get_totalprofile() public view returns (uint256) {
        return profileCount;
    }

    //------End Profile Function------//

    //------Document Function------//
    // mapping(bytes32 => Document) documents;
    bytes32[] public documentlists;
    mapping(bytes32 => Document) documents;

    // mapping(bytes32 => Document) documents;
    // bytes32[] documentlist;
    //------Signing Document------//
    function signDoc(
        bytes32 _hash_original,
        string memory _namefile,
        string memory _hash_file,
        string memory _metadata,
        string memory _hash_ipfs,
        uint256 _totalsigned,
        uint256 signing,
        uint256 _createdtime
    ) public {
        require(
            docPermission(_hash_original),
            "You are not allowed to sign this document"
        );
        if (documents[_hash_original].signing - 1 == 0) {
            //User terakhir yang melakukan ttd
            documents[_hash_original].state = 2; //State Signed
        } else {
            documents[_hash_original].state = 1; //state Process Sign
        }

        documents[_hash_original].document_id = documentCount;
        documents[_hash_original].signers = msg.sender;
        documents[_hash_original].namefile = _namefile;
        documents[_hash_original].hash_original = _hash_original;
        documents[_hash_original].hash_file = _hash_file;
        documents[_hash_original].metadata = _metadata;
        documents[_hash_original].hash_ipfs = _hash_ipfs;
        documents[_hash_original].signing = signing - 1;
        documents[_hash_original].totalsigned = _totalsigned;
        documents[_hash_original].createdtime = _createdtime;
        documents[_hash_original].completedtime = block.timestamp;
        documents[_hash_original].exist = true;
        documentlists.push(_hash_original);
        documentCount += 1;
    }

    //Add signers in ttd process
    function addsigners(
        bytes32 _hash_original,
        address _signers,
        string memory _date
    ) public {
        require(
            docPermission(_hash_original),
            "You are not allowed to sign this document"
        );
        documents[_hash_original].allsigners[_signers] = Allsigners({
            signer_address: _signers,
            state: true,
            date: _date,
            exist: true
        });
    }

    //------End Signing Document------//
    //------Get Document Latest------//

    //------End Get Document Latest------//
    //------Verify Document Function------//
    function verify_document(bytes32 _hash) public view returns (bool) {
        if (keccak256(bytes(documents[_hash].hash_file)) != keccak256(bytes(""))) {
            //Document Exist
            return true;
        }
        return false;
    }
    // ------End Document Function------//
}
